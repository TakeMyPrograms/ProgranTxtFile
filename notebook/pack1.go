package notebook
import(
"log/slog"
)
func Nop(){
	slog.Info(
	"Execution process",
	slog.String("package", "notebook"),
	slog.String("function", "Nop"),
	slog.Int("status", 200))
}
func init(){
slog.Debug("init() of notebook/notebook.go")
}