package views

import "html/template"

func NewView(files ...string) *View {
	files = append(files,
		"views/layouts/footer.gohtml",
		"views/layouts/navbar.gohtml",
		"views/layouts/bootstrap.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
	}
}

type View struct {
	Template *template.Template
}