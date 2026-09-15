<?php
// CORSヘッダーを設定
header('Access-Control-Allow-Origin: *');
header('Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS');
header('Access-Control-Allow-Headers: Content-Type');
header('Content-Type: application/json');

// プリフライトリクエスト対応
if ($_SERVER['REQUEST_METHOD'] === 'OPTIONS') {
    http_response_code(200);
    exit;
}

$request = $_SERVER['REQUEST_URI'];
$method = $_SERVER['REQUEST_METHOD'];

// ルーティング処理
if (strpos($request, '/api/users') === 0) {
    if ($method === 'GET') {
        getUsers();
    } elseif ($method === 'POST') {
        createUser();
    } elseif ($method === 'PUT') {
        updateUser();
    } elseif ($method === 'DELETE') {
        deleteUser();
    }
} elseif (strpos($request, '/api/products') === 0) {
    if ($method === 'GET') {
        getProducts();
    }
} elseif (strpos($request, '/api/health') === 0) {
    health();
} else {
    http_response_code(404);
    echo json_encode(['error' => 'Not Found']);
}

// ユーザー一覧を取得
function getUsers() {
    // 実際にはデータベースから取得
    $users = [
        ['id' => 1, 'name' => 'Alice', 'email' => 'alice@example.com'],
        ['id' => 2, 'name' => 'Bob', 'email' => 'bob@example.com'],
        ['id' => 3, 'name' => 'Charlie', 'email' => 'charlie@example.com'],
    ];
    echo json_encode(['status' => 'success', 'data' => $users]);
}

// ユーザーを作成
function createUser() {
    $input = json_decode(file_get_contents('php://input'), true);
    
    if (!isset($input['name']) || !isset($input['email'])) {
        http_response_code(400);
        echo json_encode(['error' => 'name and email are required']);
        return;
    }
    
    // 実際にはデータベースに保存
    $newUser = [
        'id' => 4,
        'name' => $input['name'],
        'email' => $input['email']
    ];
    
    http_response_code(201);
    echo json_encode(['status' => 'success', 'data' => $newUser]);
}

// ユーザーを更新
function updateUser() {
    $input = json_decode(file_get_contents('php://input'), true);
    
    // 実際にはデータベースを更新
    echo json_encode(['status' => 'success', 'message' => 'User updated']);
}

// ユーザーを削除
function deleteUser() {
    echo json_encode(['status' => 'success', 'message' => 'User deleted']);
}

// 商品一覧を取得
function getProducts() {
    $products = [
        ['id' => 1, 'name' => 'Laptop', 'price' => 1000],
        ['id' => 2, 'name' => 'Mouse', 'price' => 25],
        ['id' => 3, 'name' => 'Keyboard', 'price' => 75],
    ];
    echo json_encode(['status' => 'success', 'data' => $products]);
}

// ヘルスチェック
function health() {
    echo json_encode(['status' => 'ok']);
}
?>
