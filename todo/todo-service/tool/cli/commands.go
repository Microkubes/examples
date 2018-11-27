// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "microtodo": CLI Commands
//
// Command:
// $ goagen
// --design=github.com/Microkubes/examples/todo/todo-service/design
// --out=$(GOPATH)src/github.com/Microkubes/examples/todo/todo-service
// --version=v1.3.1

package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Microkubes/examples/todo/todo-service/client"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/spf13/cobra"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	// AddTodoTodoCommand is the command line data structure for the addTodo action of todo
	AddTodoTodoCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// DeleteTodoTodoCommand is the command line data structure for the deleteTodo action of todo
	DeleteTodoTodoCommand struct {
		// Todo ID
		TodoID      string
		PrettyPrint bool
	}

	// FilterTodosTodoCommand is the command line data structure for the filterTodos action of todo
	FilterTodosTodoCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// GetAllTodosTodoCommand is the command line data structure for the getAllTodos action of todo
	GetAllTodosTodoCommand struct {
		// Limit todos per page
		Limit int
		// number of todos to skip
		Offset int
		// order by
		Order       string
		Sorting     string
		PrettyPrint bool
	}

	// GetByIDTodoCommand is the command line data structure for the getById action of todo
	GetByIDTodoCommand struct {
		// Todo ID
		TodoID      string
		PrettyPrint bool
	}

	// UpdateTodoTodoCommand is the command line data structure for the updateTodo action of todo
	UpdateTodoTodoCommand struct {
		Payload     string
		ContentType string
		// Todo ID
		TodoID      string
		PrettyPrint bool
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "add-todo",
		Short: `Add new todo`,
	}
	tmp1 := new(AddTodoTodoCommand)
	sub = &cobra.Command{
		Use:   `todo ["/todo/add"]`,
		Short: ``,
		Long: `

Payload example:

{
   "description": "Sed minus.",
   "title": "Et reprehenderit officia aut."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "delete-todo",
		Short: `Delete todo`,
	}
	tmp2 := new(DeleteTodoTodoCommand)
	sub = &cobra.Command{
		Use:   `todo ["/todo/TODOID/delete"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "filter-todos",
		Short: `Filter (lookup) todos`,
	}
	tmp3 := new(FilterTodosTodoCommand)
	sub = &cobra.Command{
		Use:   `todo ["/todo/filter"]`,
		Short: ``,
		Long: `

Payload example:

{
   "filter": "1971-12-03T04:24:10Z",
   "order": [
      {
         "direction": "Placeat fugiat mollitia reiciendis.",
         "property": "Nostrum accusantium molestias blanditiis nam."
      },
      {
         "direction": "Placeat fugiat mollitia reiciendis.",
         "property": "Nostrum accusantium molestias blanditiis nam."
      }
   ],
   "page": 3307553241945421545,
   "pageSize": 2433291310988242551
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "get-all-todos",
		Short: `Get all todos`,
	}
	tmp4 := new(GetAllTodosTodoCommand)
	sub = &cobra.Command{
		Use:   `todo ["/todo/all"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "get-byid",
		Short: `Get todo by ID`,
	}
	tmp5 := new(GetByIDTodoCommand)
	sub = &cobra.Command{
		Use:   `todo ["/todo/TODOID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "update-todo",
		Short: `Update todo`,
	}
	tmp6 := new(UpdateTodoTodoCommand)
	sub = &cobra.Command{
		Use:   `todo [("/todo/TODOID"|"/todo/TODOID")]`,
		Short: ``,
		Long: `

Payload example:

{
   "description": "Eum dolores dolore similique soluta officiis.",
   "done": false,
   "title": "Fuga soluta."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp6.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run makes the HTTP request corresponding to the AddTodoTodoCommand command.
func (cmd *AddTodoTodoCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/todo/add"
	}
	var payload client.TodoPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.AddTodoTodo(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *AddTodoTodoCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the DeleteTodoTodoCommand command.
func (cmd *DeleteTodoTodoCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/todo/%v/delete", url.QueryEscape(cmd.TodoID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteTodoTodo(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteTodoTodoCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var todoID string
	cc.Flags().StringVar(&cmd.TodoID, "todoID", todoID, `Todo ID`)
}

// Run makes the HTTP request corresponding to the FilterTodosTodoCommand command.
func (cmd *FilterTodosTodoCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/todo/filter"
	}
	var payload client.FilterTodoPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.FilterTodosTodo(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *FilterTodosTodoCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the GetAllTodosTodoCommand command.
func (cmd *GetAllTodosTodoCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/todo/all"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.GetAllTodosTodo(ctx, path, intFlagVal("limit", cmd.Limit), intFlagVal("offset", cmd.Offset), stringFlagVal("order", cmd.Order), stringFlagVal("sorting", cmd.Sorting))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *GetAllTodosTodoCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var limit int
	cc.Flags().IntVar(&cmd.Limit, "limit", limit, `Limit todos per page`)
	var offset int
	cc.Flags().IntVar(&cmd.Offset, "offset", offset, `number of todos to skip`)
	var order string
	cc.Flags().StringVar(&cmd.Order, "order", order, `order by`)
	var sorting string
	cc.Flags().StringVar(&cmd.Sorting, "sorting", sorting, ``)
}

// Run makes the HTTP request corresponding to the GetByIDTodoCommand command.
func (cmd *GetByIDTodoCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/todo/%v", url.QueryEscape(cmd.TodoID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.GetByIDTodo(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *GetByIDTodoCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var todoID string
	cc.Flags().StringVar(&cmd.TodoID, "todoID", todoID, `Todo ID`)
}

// Run makes the HTTP request corresponding to the UpdateTodoTodoCommand command.
func (cmd *UpdateTodoTodoCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/todo/%v", url.QueryEscape(cmd.TodoID))
	}
	var payload client.TodoUpdatePayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateTodoTodo(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateTodoTodoCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var todoID string
	cc.Flags().StringVar(&cmd.TodoID, "todoID", todoID, `Todo ID`)
}
