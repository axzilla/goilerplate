// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "fmt"
import "github.com/axzilla/goilerplate/pkg/icons"

type ToastProps struct {
	Message     string // Die Nachricht
	Type        string // success, error, warning, info, default
	Position    string // top-right, top-left, top-center, bottom-right, bottom-left, bottom-center
	Duration    int    // Millisekunden, 0 = permanent
	Dismissible bool   // Kann manuell geschlossen werden
	Size        string // sm, md, lg
	Icon        bool   // Icon anzeigen/verstecken
	HTML        string // Custom HTML Content
}

func Toast(cfg ToastProps) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-data=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(`{
            show: true,
            message: '` + cfg.Message + `',
            type: '` + cfg.Type + `',
            position: '` + cfg.Position + `',
            duration: ` + fmt.Sprint(cfg.Duration) + `,
            dismissible: ` + fmt.Sprint(cfg.Dismissible) + `,
            size: '` + cfg.Size + `'
        }`)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/toast.templ`, Line: 27, Col: 10}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" x-init=\"if(duration &gt; 0) setTimeout(() =&gt; show = false, duration)\" x-show=\"show\" x-transition:enter=\"transition ease-out duration-300\" x-transition:enter-start=\"opacity-0 translate-y-4\" x-transition:enter-end=\"opacity-100 translate-y-0\" x-transition:leave=\"transition ease-in duration-200\" x-transition:leave-start=\"opacity-100 translate-y-0\" x-transition:leave-end=\"opacity-0 translate-y-4\" @click=\"if(dismissible) show = false\" class=\"fixed pointer-events-auto\" :class=\"{\n            // Position\n            &#39;top-4 right-4&#39;: position === &#39;top-right&#39;,\n            &#39;top-4 left-4&#39;: position === &#39;top-left&#39;,\n            &#39;top-4 left-1/2 -translate-x-1/2&#39;: position === &#39;top-center&#39;,\n            &#39;bottom-4 right-4&#39;: position === &#39;bottom-right&#39;,\n            &#39;bottom-4 left-4&#39;: position === &#39;bottom-left&#39;,\n            &#39;bottom-4 left-1/2 -translate-x-1/2&#39;: position === &#39;bottom-center&#39;,\n            // Size\n            &#39;w-72&#39;: size === &#39;sm&#39;,\n            &#39;w-96&#39;: size === &#39;md&#39;,\n            &#39;w-[30rem]&#39;: size === &#39;lg&#39;\n        }\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if cfg.HTML != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-html=\"message\" class=\"rounded-lg shadow-lg\"></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <div class=\"bg-primary-foreground rounded-lg shadow-sm border p-4 flex items-center justify-center\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if cfg.Icon {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if cfg.Type == "success" {
					templ_7745c5c3_Err = icons.CircleCheck(icons.IconProps{Size: "18", Class: "text-green-500 mr-3"}).Render(ctx, templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if cfg.Type == "error" {
					templ_7745c5c3_Err = icons.CircleX(icons.IconProps{Size: "18", Class: "text-red-500 mr-3"}).Render(ctx, templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if cfg.Type == "warning" {
					templ_7745c5c3_Err = icons.TriangleAlert(icons.IconProps{Size: "18", Class: "text-yellow-500 mr-3"}).Render(ctx, templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if cfg.Type == "info" {
					templ_7745c5c3_Err = icons.Info(icons.IconProps{Size: "18", Class: "text-blue-500 mr-3"}).Render(ctx, templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex-1\" x-text=\"message\"></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if cfg.Dismissible {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button @click.stop=\"show = false\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = icons.X(icons.IconProps{
					Size:  "18",
					Class: "ml-4 flex-shrink-0 opacity-75 hover:opacity-100",
				}).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate