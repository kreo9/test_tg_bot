package product

var allProducts = []Product{
	{Title: "Крем"},
	{Title: "Соль для ванной"},
	{Title: "Мыло"},
	{Title: "Шампунь"},
	{Title: "Умывалка"},
}

type Product struct {
	Title string
}
