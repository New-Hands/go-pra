package ui

// A simple program demonstrating the text area component from the Bubbles
// component library.
import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"strconv"
	"strings"
	"time"
)

type NetInMsg string

type (
	errMsg error
)

// 控制和数据分离 框架调度model
// a model that describes the application state and three simple methods on that model:
// Init, a function that returns an initial command for the application to run.
// Update, a function that handles incoming events and updates the model accordingly.
// View, a function that renders the UI based on the data in the model.

type Model struct {
	viewport     viewport.Model
	messages     []string
	textarea     textarea.Model
	spinner      spinner.Model
	senderStyle  lipgloss.Style
	errorStyle   lipgloss.Style
	receiveStyle lipgloss.Style
	err          error
	inputHandle  func(input string) (string, error)
}

// InitialModel m copy
func InitialModel(handle func(input string) (string, error)) Model {
	profile := termenv.DefaultOutput().ColorProfile()

	ta := textarea.New()
	ta.Placeholder = "Send a message..."
	ta.Focus()
	ta.Prompt = "┃ "
	ta.CharLimit = 20480
	ta.SetWidth(100)
	ta.SetHeight(2)
	// Remove cursor line styling
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
	ta.ShowLineNumbers = false
	ta.KeyMap.InsertNewline.SetEnabled(false)

	vp := viewport.New(30, 10)
	// 修改keymap
	vp.KeyMap = viewport.KeyMap{
		PageDown: key.NewBinding(
			key.WithKeys("pgdown"),
			key.WithHelp("pgdn", "page down"),
		),
		PageUp: key.NewBinding(
			key.WithKeys("pgup"),
			key.WithHelp("pgup", "page up"),
		),
		HalfPageUp: key.NewBinding(
			key.WithKeys("ctrl+u"),
			key.WithHelp("ctrl+u", "half page up"),
		),
		HalfPageDown: key.NewBinding(
			key.WithKeys("ctrl+d"),
			key.WithHelp("ctrl+d", "half page down"),
		),
		Up: key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("↑", "up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("↓", "down"),
		),
	}

	sp := spinner.Spinner{
		Frames: []string{".-..-..---..---.   .---..-.   .-.\n| .` || |- `| |'   | |  | |__ | |\n`-'`-'`---' `-'    `---'`----'`-'\nprofile " + strconv.Itoa(int(profile))},
		FPS:    time.Minute,
	}

	withSpinner := spinner.WithSpinner(sp)
	style := spinner.WithStyle(lipgloss.NewStyle().Bold(true).Italic(true).
		Height(4).Foreground(lipgloss.Color("2")))
	spModel := spinner.New(withSpinner, style)

	return Model{
		textarea:     ta,
		messages:     []string{},
		viewport:     vp,
		spinner:      spModel,
		senderStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		receiveStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("3")),
		errorStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("1")),
		err:          nil,
		inputHandle:  handle,
	}
}

func (m Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
		spCmd tea.Cmd
	)

	// update child
	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)
	m.spinner, spCmd = m.spinner.Update(msg)

	// 检查msg大小内存中只保存前50条
	mLen := len(m.messages)
	if mLen > 50 {
		m.messages = m.messages[mLen-50:]
	}

	tWidth := m.viewport.Width
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			fmt.Println(m.textarea.Value())
			return m, tea.Quit
		case tea.KeyEnter:
			// 数据处理
			pInput, err := m.inputHandle(m.textarea.Value())
			if err != nil {
				// 替换添加样式
				elems := m.errorStyle.Render("error:")
				newMsg := splitMsg(tWidth, "error:"+err.Error())
				newMsg[0] = strings.Replace(newMsg[0], "error", elems, -1)
				appendMsg(&m.messages, newMsg)
			} else {
				// message addr change
				elems := m.senderStyle.Render("send:")
				newMsg := splitMsg(tWidth, "send:"+pInput)
				newMsg[0] = strings.Replace(newMsg[0], "send:", elems, -1)
				appendMsg(&m.messages, newMsg)
			}
			m.viewport.SetContent(strings.Join(m.messages, "\n"))
			m.textarea.Reset()
			m.viewport.GotoBottom()
		}

	case tea.WindowSizeMsg:
		// different code style
		m.textarea.SetWidth(msg.Width)
		m.viewport.Width = msg.Width
		minH := 5

		otherH := m.textarea.Height() + m.spinner.Style.GetHeight()
		if (msg.Height - (otherH + 4)) > minH {
			// 减去 textArea
			minH = msg.Height - otherH - 4

		}
		m.viewport.Height = minH

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		elems := m.errorStyle.Render("error:")
		newMsg := splitMsg(tWidth, "error:"+msg.Error())
		newMsg[0] = strings.Replace(newMsg[0], "error:", elems, -1)
		appendMsg(&m.messages, newMsg)

		m.viewport.SetContent(strings.Join(m.messages, "\n"))
		m.viewport.GotoBottom()
		return m, nil
	case NetInMsg:
		elems := m.receiveStyle.Render("receive:")
		newMsg := splitMsg(tWidth, "receive:"+string(msg))
		newMsg[0] = strings.Replace(newMsg[0], "receive:", elems, -1)
		appendMsg(&m.messages, newMsg)
		m.viewport.SetContent(strings.Join(m.messages, "\n"))
		m.viewport.GotoBottom()
	}

	return m, tea.Batch(tiCmd, vpCmd, spCmd)
}

func appendMsg(oriMsg *[]string, newMsg []string) {
	for i := 0; i < len(newMsg); i++ {
		*oriMsg = append(*oriMsg, newMsg[i])
	}
}

// 分割消息
func splitMsg(widthLimit int, strMsg string) []string {
	mLen := len(strMsg)

	var msg []string

	for true {
		last := mLen - widthLimit
		if last <= 0 {
			return append(msg, strMsg)
		} else {
			msg = append(msg, strMsg[:widthLimit])
			// 改变消息
			strMsg = strMsg[widthLimit:]
			mLen = last
		}
	}

	return msg
}

// View 网络面板视图渲染数据
func (m Model) View() string {
	view := m.spinner.View()
	return fmt.Sprintf(
		"%s\n%s\n\n%s",
		view,
		m.viewport.View(),
		m.textarea.View(),
	) + "\n\n"
}
