package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const PHPBackend = "http://localhost:8080"

// User構造体
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Product構造体
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// APIレスポンス
type APIResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
}

func main() {
	// ホームページ
	http.HandleFunc("/", handleHome)
	
	// ユーザーページ
	http.HandleFunc("/users", handleUsers)
	
	// ユーザー追加ページ
	http.HandleFunc("/add-user", handleAddUser)
	
	// 商品ページ
	http.HandleFunc("/products", handleProducts)
	
	// API プロキシ
	http.HandleFunc("/api/", handleAPIProxy)
	
	fmt.Println("Go frontend server running on :8000")
	http.ListenAndServe(":8000", nil)
}

// ホームページ
func handleHome(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Go + PHP デモ</title>
		<style>
			body { font-family: Arial; margin: 20px; background: #f5f5f5; }
			h1 { color: #333; }
			a { display: inline-block; margin: 10px 5px 0 0; }
			button { padding: 10px 20px; font-size: 16px; background: #007bff; color: white; border: none; cursor: pointer; border-radius: 4px; }
			button:hover { background: #0056b3; }
		</style>
	</head>
	<body>
		<h1>Go + PHP デモアプリケーション</h1>
		<p>フロントエンド: Go | バックエンド: PHP</p>
		<div>
			<a href="/users"><button>ユーザー一覧</button></a>
			<a href="/add-user"><button>ユーザー追加</button></a>
			<a href="/products"><button>商品一覧</button></a>
		</div>
	</body>
	</html>
	`
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprint(w, html)
}

// ユーザー一覧ページ
func handleUsers(w http.ResponseWriter, r *http.Request) {
	// PHPバックエンドからユーザー情報を取得
	resp, err := http.Get(PHPBackend + "/api/users")
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var apiResp APIResponse
	json.Unmarshal(body, &apiResp)

	users := apiResp.Data.([]interface{})

	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>ユーザー一覧</title>
		<style>
			body { font-family: Arial; margin: 20px; }
			table { border-collapse: collapse; width: 100%; }
			th, td { border: 1px solid #ddd; padding: 12px; text-align: left; }
			th { background: #007bff; color: white; }
			tr:nth-child(even) { background: #f9f9f9; }
			button { padding: 8px 16px; cursor: pointer; margin-top: 20px; }
		</style>
	</head>
	<body>
		<h1>ユーザー一覧</h1>
		<table>
			<tr>
				<th>ID</th>
				<th>名前</th>
				<th>メール</th>
			</tr>
	`

	for _, u := range users {
		user := u.(map[string]interface{})
		html += fmt.Sprintf(`
			<tr>
				<td>%v</td>
				<td>%v</td>
				<td>%v</td>
			</tr>
		`, user["id"], user["name"], user["email"])
	}

	html += `
		</table>
		<a href="/"><button>戻る</button></a>
	</body>
	</html>
	`
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprint(w, html)
}

// ユーザー追加ページ
func handleAddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		name := r.FormValue("name")
		email := r.FormValue("email")

		// PHPバックエンドにPOSTリクエスト
		data := fmt.Sprintf(`{"name": "%s", "email": "%s"}`, name, email)
		resp, err := http.Post(
			PHPBackend+"/api/users",
			"application/json",
			strings.NewReader(data),
		)

		if err == nil {
			resp.Body.Close()
			html := `
			<!DOCTYPE html>
			<html>
			<head>
				<title>完了</title>
				<style>body { font-family: Arial; margin: 20px; }</style>
			</head>
			<body>
				<h1 style="color: green;">✓ ユーザーを追加しました！</h1>
				<a href="/users"><button>ユーザー一覧を表示</button></a>
				<a href="/"><button>ホームに戻る</button></a>
			</body>
			</html>
			`
			w.Header().Set("Content-Type", "text/html; charset=UTF-8")
			fmt.Fprint(w, html)
			return
		}
	}

	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>ユーザー追加</title>
		<style>
			body { font-family: Arial; margin: 20px; }
			form { max-width: 400px; }
			input { display: block; margin: 10px 0; padding: 8px; width: 100%; box-sizing: border-box; }
			button { padding: 10px 20px; font-size: 16px; background: #28a745; color: white; border: none; cursor: pointer; }
		</style>
	</head>
	<body>
		<h1>新規ユーザー追加</h1>
		<form method="POST">
			<input type="text" name="name" placeholder="名前" required>
			<input type="email" name="email" placeholder="メールアドレス" required>
			<button type="submit">追加</button>
		</form>
		<a href="/"><button style="background: #6c757d; margin-top: 10px;">戻る</button></a>
	</body>
	</html>
	`
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprint(w, html)
}

// 商品一覧ページ
func handleProducts(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(PHPBackend + "/api/products")
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var apiResp APIResponse
	json.Unmarshal(body, &apiResp)

	products := apiResp.Data.([]interface{})

	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>商品一覧</title>
		<style>
			body { font-family: Arial; margin: 20px; }
			.products { display: grid; grid-template-columns: repeat(auto-fill, minmax(250px, 1fr)); gap: 20px; }
			.product { border: 1px solid #ddd; padding: 15px; border-radius: 4px; background: #fff; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
			.product h3 { margin-top: 0; }
			.price { font-size: 20px; color: #28a745; font-weight: bold; }
			button { padding: 8px 16px; margin-top: 20px; cursor: pointer; }
		</style>
	</head>
	<body>
		<h1>商品一覧</h1>
		<div class="products">
	`

	for _, p := range products {
		product := p.(map[string]interface{})
		html += fmt.Sprintf(`
			<div class="product">
				<h3>%v</h3>
				<p class="price">$%v</p>
			</div>
		`, product["name"], product["price"])
	}

	html += `
		</div>
		<a href="/"><button>戻る</button></a>
	</body>
	</html>
	`
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprint(w, html)
}

// APIプロキシ（フロントエンドがAPI呼び出しを仲介する場合）
func handleAPIProxy(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	resp, err := http.Get(PHPBackend + path)
	if err != nil {
		http.Error(w, "Backend error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, resp.Body)
}
