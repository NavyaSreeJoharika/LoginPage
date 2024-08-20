/*
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	// Serve the login HTML and CSS files
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "login.html")
	})
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "style.css")
	})
	http.HandleFunc("/greeting.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "greeting.html")
	})

	// Handle the form submission
	http.HandleFunc("/login", loginHandler)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	reader := bufio.NewReader(r.Body)
	data, err := reader.ReadString('\n')
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// Parse the form data
	formData, err := url.ParseQuery(data)
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusInternalServerError)
		return
	}

	name := formData.Get("name")
	phone := formData.Get("phone")

	// Basic validation (you can expand this as needed)
	if name == "" || phone == "" {
		http.Error(w, "Name and phone number are required", http.StatusBadRequest)
		return
	}

	// Redirect to thank you page
	http.Redirect(w, r, "/thankyou", http.StatusSeeOther)
}

func thankYouHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Thank you for login")
}
*/
/*
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

func main() {
	// Serve the login HTML and CSS files
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "login.html")
	})
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "style.css")
	})
	http.HandleFunc("/greeting.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "greeting.html")
	})

	// Handle the form submission
	http.HandleFunc("/login", loginHandler)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")

	// Redirect to the greeting page with the name as a query parameter
	http.Redirect(w, r, "/greet?name="+url.QueryEscape(name), http.StatusSeeOther)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	// Load the greeting template
	tmpl, err := template.ParseFiles("greeting.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	// Inject the name into the template
	err = tmpl.Execute(w, map[string]string{"Name": name})
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}
*/
/*
package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

func main() {
	// Serve the login HTML and CSS files
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "login.html")
	})
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "style.css")
	})

	// Handle the form submission
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/greet", greetHandler)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	reader := bufio.NewReader(r.Body)
	data, err := reader.ReadString('\n')
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// Parse the form data
	formData, err := url.ParseQuery(data)
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusInternalServerError)
		return
	}

	name := formData.Get("name")

	// Redirect to the greeting page with the name as a query parameter
	http.Redirect(w, r, "/greet?name="+url.QueryEscape(name), http.StatusSeeOther)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	// Load the greeting template
	tmpl, err := template.ParseFiles("greeting.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	// Inject the name into the template
	err = tmpl.Execute(w, map[string]string{"Name": name})
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}
*/
/*
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

func main() {
	// Serve the login HTML and CSS files
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "login.html")
	})
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "style.css")
	})

	// Handle the form submission
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/greet", greetHandler)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")

	// Redirect to the greeting page with the name as a query parameter
	http.Redirect(w, r, "/greet?name="+url.QueryEscape(name), http.StatusSeeOther)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	// Load the greeting template
	tmpl, err := template.ParseFiles("greeting.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	// Inject the name into the template
	err = tmpl.Execute(w, map[string]string{"Name": name})
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}
*/

package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

func main() {
	// Serve the login HTML and CSS files
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "login.html")
	})
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "style.css")
	})

	// Handle the form submission
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/greet", greetHandler)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var buf bytes.Buffer
	reader := bufio.NewReader(r.Body)
	_, err := buf.ReadFrom(reader)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	formData, err := url.ParseQuery(buf.String())
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusInternalServerError)
		return
	}

	// name := formData.Get("name")
	// number := formData.Get("number")
	// address := formData.Get("address")
	// Redirect to the greeting page with the name as a query parameter
	//http.Redirect(w, r, "/greet?name="+url.QueryEscape(name)+"&number="+url.QueryEscape(number)+"&address="+url.QueryEscape(address), http.StatusSeeOther)
	r = r.WithContext(context.WithValue(r.Context(), "formData", formData))
	greetHandler(w, r)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	formData := r.Context().Value("formData").(url.Values)
	name := formData.Get("name")
	number := formData.Get("number")
	address := formData.Get("address")
	// Load the greeting template
	tmpl, err := template.ParseFiles("greeting.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	// Inject the name into the template
	err = tmpl.Execute(w, map[string]string{"Name": name, "Number": number, "Address": address})
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}

}
