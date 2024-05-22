// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/saintmalik/delta/model"

// import "fmt"
func Dash(packages []model.Mypack) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.Raw("<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script src=\"https://unpkg.com/htmx.org@1.9.12\" type=\"text/javascript\"></script><script src=\"https://cdn.tailwindcss.com\"></script><title>Package Updates</title></head><body><div class=\"flex flex-col h-screen\"><div class=\"px-4 py-4 md:px-6\"><div class=\"flex flex-col items-start gap-4 md:flex-row md:items-center md:justify-between\"><div class=\"flex items-center gap-4\"><h1 class=\"text-2xl font-bold\">Package Dashboard</h1></div></div></div><div class=\"flex justify-end mt-4 gap-2 p-4 md:p-6\"><a class=\"inline-flex h-9 items-center justify-center rounded-md bg-gray-900 px-4 py-2 text-sm font-medium text-gray-50 shadow transition-colors dark:border-gray-800 hover:bg-gray-900/90 focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-gray-950 disabled:pointer-events-none disabled:opacity-50 dark:bg-gray-50 dark:text-gray-900 dark:hover:bg-gray-50/90 dark:focus-visible:ring-gray-300\" hx-swap=\"transition:true\" href=\"/check\">Check Changelogs</a> <a class=\"inline-flex h-9 items-center justify-center rounded-md border border-gray-200 bg-gray-700 px-4 py-2 text-sm font-medium shadow-sm transition-colors hover:bg-gray-600 hover:text-white focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-gray-950 disabled:pointer-events-none disabled:opacity-50 dark:border-gray-800 dark:bg-gray-700 dark:text-white dark:hover:bg-gray-600 dark:focus-visible:ring-gray-300\" hx-swap=\"transition:true\" href=\"/package\">Add Package</a></div><div class=\"flex-1 overflow-auto p-4 md:p-6\"><div class=\"border rounded-lg shadow-sm overflow-x-auto\"><div class=\"relative w-full overflow-auto\"><table class=\"w-full caption-bottom text-sm\"><thead class=\"[&amp;_tr]:border-b\"><tr class=\"border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted\"><th class=\"h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0\">Package Name</th><th class=\"h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0\">Package Version</th><th class=\"h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0\">Package URL</th></tr></thead> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(packages) != 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<tbody class=\"[&amp;_tr:last-child]:border-0\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, pack := range packages {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<tr class=\"border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted\"><td class=\"p-4 align-middle [&amp;:has([role=checkbox])]:pr-0\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var2 string
				templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(pack.PackageName)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/dash.templ`, Line: 65, Col: 95}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td class=\"p-4 align-middle [&amp;:has([role=checkbox])]:pr-0\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(pack.PackageVersion)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/dash.templ`, Line: 66, Col: 98}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td class=\"p-4 align-middle [&amp;:has([role=checkbox])]:pr-0\"><a hx-swap=\"transition:true\" href=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 templ.SafeURL = templ.URL(pack.PackageURL)
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var4)))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"badge badge-primary p-3 hover:scale-[1.1] text-blue-500\" target=\"_blank\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(pack.PackageName)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/dash.templ`, Line: 73, Col: 36}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></td></tr>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</tbody>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<tbody><tr><td colspan=\"4\" align=\"center\">You do not have any package yet </td></tr></tbody>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</table></div></div></div></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}