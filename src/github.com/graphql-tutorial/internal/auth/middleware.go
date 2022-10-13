package auth

import (
	"context"
	"net/http"
	"strconv"

	"github.com/graphql-tutorial/internal/users"
	"github.com/graphql-tutorial/pkg/jwt"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// 未認証のユーザを許可する
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			// jwtトークンの有効化
			tokenStr := header
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// ユーザを作成し、データベース内にユーザが存在するかどうかをチェックする
			user := users.User{Username: username}
			id, err := users.GetUserIdByUsername(username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			user.ID = strconv.Itoa(id)
			// コンテキストに入れる
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			// そして、新しいコンテキストで次を呼び出す
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext: コンテキストからユーザを見つける。ミドルウェアが動作している必要がある。
func ForContext(ctx context.Context) *users.User {
	raw, _ := ctx.Value(userCtxKey).(*users.User)
	return raw
}
