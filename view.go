// templates
// parse and cache templates
var templates = template.Must(template.ParseFiles(
    "templates/components/login/login-element.tmpl",
    "templates/components/login/login-cdn-library.tmpl",
    "templates/components/login/login-script.tmpl",
    "templates/components/login/login-return-script.tmpl",
    "pages/login.tmpl",
    "pages/home.tmpl",
))