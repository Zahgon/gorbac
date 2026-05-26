package main

import (
	"database/sql"
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"sync"

	gorbac "github.com/mikespook/gorbac/v3"
	"github.com/mikespook/possum"
	_ "modernc.org/sqlite"
)

//go:embed static
var staticFS embed.FS

// A permission is always "resource::action" — layer type, sep "::".
// We store only the composed ID; type/sep are implicit.
type roleState struct {
	parents     []string
	permissions map[string]struct{} // key = "resource::action"
}

type store struct {
	mu    sync.RWMutex
	db    *sql.DB
	roles map[string]*roleState
	rbac  *gorbac.RBAC[string]
}

var state *store

// ---- DB ----

const schema = `
CREATE TABLE IF NOT EXISTS roles (id TEXT PRIMARY KEY);
CREATE TABLE IF NOT EXISTS role_parents (
    role_id TEXT NOT NULL, parent_id TEXT NOT NULL,
    PRIMARY KEY (role_id, parent_id)
);
CREATE TABLE IF NOT EXISTS role_permissions (
    role_id TEXT NOT NULL, perm_id TEXT NOT NULL,
    PRIMARY KEY (role_id, perm_id)
);`

func openDB(path string) (*sql.DB, bool, error) { _ = "STUB: not implemented"; return nil, false, nil }

func loadFromDB(db *sql.DB) (*store, error) { _ = "STUB: not implemented"; return nil, nil }

// rebuildRBAC must be called with mu.Lock held (or single-threaded init).
func (s *store) rebuildRBAC() { _ = "STUB: not implemented"; return }

// ---- Assertions ----

func buildAssertion(mode string, params map[string]string) gorbac.AssertionFunc[string] {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec

// ---- Wire types ----

type RoleView struct {
	ID          string   `json:"id"`
	Parents     []string `json:"parents"`
	Permissions []string `json:"permissions"` // "resource::action" strings
}

type StateResponse struct {
	Roles []RoleView `json:"roles"`
}

// ---- Handlers ----

func handleState(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func handleRoles(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func handleRoleByID(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func deleteRole(w http.ResponseWriter, r *http.Request, roleID string) {
	_ = "STUB: not implemented"
	return
}

func assignPermission(w http.ResponseWriter, r *http.Request, roleID string) {
	_ = "STUB: not implemented"
	return
}

// "resource::action"

func revokePermission(w http.ResponseWriter, r *http.Request, roleID, permID string) {
	_ = "STUB: not implemented"
	return
}

type verifyReq struct {
	Role         string            `json:"role"`
	Resource     string            `json:"resource"` // combined with action → "resource::action"
	Action       string            `json:"action"`
	CheckMode    string            `json:"checkMode"` // "single" | "any" | "all"
	AssertMode   string            `json:"assertMode"`
	AssertParams map[string]string `json:"assertParams"`
	Roles        []string          `json:"roles"` // extra roles for any/all
}

func handleVerify(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// ---- Seed ----

var seedData = struct {
	roles []struct {
		id      string
		parents []string
	}
	assigns []struct{ role, perm string }
}{
	roles: []struct {
		id      string
		parents []string
	}{
		{"guest", nil},
		{"member", []string{"guest"}},
		{"editor", []string{"member"}},
		{"admin", []string{"editor"}},
		{"superadmin", []string{"admin"}},
	},
	assigns: []struct{ role, perm string }{
		{"guest", "article::read"},
		{"guest", "comment::read"},
		{"guest", "profile::read"},
		{"guest", "image::read"},
		{"guest", "video::read"},
		{"member", "comment::edit"},
		{"member", "profile::edit"},
		{"editor", "article::edit"},
		{"editor", "article::delete"},
		{"editor", "image::edit"},
		{"editor", "image::delete"},
		{"editor", "video::edit"},
		{"editor", "video::delete"},
		{"editor", "tag::edit"},
		{"editor", "category::edit"},
		{"admin", "report::edit"},
		{"admin", "setting::edit"},
		{"admin", "comment::delete"},
		{"superadmin", "audit-log::read"},
		{"superadmin", "profile::delete"},
	},
}

func applySeed(db *sql.DB, s *store) error { _ = "STUB: not implemented"; return nil }

func handleReset(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func main() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "playground.db"
	}

	db, firstCreate, err := openDB(dbPath)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()

	state, err = loadFromDB(db)
	if err != nil {
		log.Fatalf("load db: %v", err)
	}
	if firstCreate {
		if err = applySeed(db, state); err != nil {
			log.Fatalf("seed db: %v", err)
		}
	}

	sub, err := fs.Sub(staticFS, "static")
	if err != nil {
		log.Fatal(err)
	}

	chain := func(h http.HandlerFunc, methods ...string) http.HandlerFunc {
		return possum.Chain(h, possum.Log, possum.Cors(nil), possum.AllowMethods(methods...))
	}

	http.HandleFunc("/api/state", chain(handleState, "GET"))
	http.HandleFunc("/api/roles", chain(handleRoles, "POST"))
	http.HandleFunc("/api/roles/", chain(handleRoleByID, "DELETE", "POST"))
	http.HandleFunc("/api/seed", chain(handleReset, "POST"))
	http.HandleFunc("/api/verify", chain(handleVerify, "POST"))
	http.Handle("/", http.FileServer(http.FS(sub)))

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":9080"
	}
	log.Printf("RBAC Playground listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
