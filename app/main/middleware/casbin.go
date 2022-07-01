package middleware

import (
	"log"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
)

func RBACMiddleware(obj string, act string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			e, err := casbin.NewEnforcer("./casbin_model.conf", "./casbin_policy.csv")
			if err != nil {
				log.Fatalf("unable to create casbin enforcer: %v", err)
			}

			sub, _ := c.Get("role").(string)

			a := e.GetPolicy()
			rolesMap := make(map[string]bool)
			for _, v := range a {
				if len(v) > 0 {
					rolesMap[v[0]] = true
				}
			}

			res, err := e.Enforce(sub, obj, act)
			if !res || err != nil {
				return c.JSON(http.StatusForbidden, responseForbidden)
			}

			return next(c)
		}
	}
}
