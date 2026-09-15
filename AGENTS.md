# Hospital

Full-stack hospital management app. The real code is the Go backend; the frontend is a Next.js dashboard starter kit that is not yet wired to the API.

## Layout

- `backend/` — Go API: Gin + GORM + PostgreSQL. Single entrypoint `backend/cmd/api/main.go`; all dependency wiring lives there.
- `backend/internal/<context>/{domain,application,infrastructure}` — clean architecture, bounded contexts: `users`, `patients`, `doctors`, `specialty`.
- `backend/Hospital/` — Bruno API collection (`.yml` files; `opencollection.yml`).
- `frontend/` — Next.js 15 (pages router) dashboard template (Xintra "starterkit"), with Redux and Firebase deps. No API client for the Go backend exists yet.
- `architecture-runtime.html` — generated diagram artifact (untracked).

## Backend (run from `backend/`)

- `make run` — start API on `:8080`; loads `.env` via godotenv if present.
- `make db-up` / `make db-down` — start/stop only Postgres via `build/docker-compose.yml`.
- `make docker-up` — full stack (Postgres + API) on `:8080`.
- `make test` — `go test ./... -v`. Note: there are currently **no `_test.go` files**, so this passes vacuously.
- **`make seed` is broken**: it references `scripts/seed.sh`, which doesn't exist.
- Config comes from env vars (defaults in `internal/shared/infrastructure/config/config.go`). Copy `.env.sample` → `.env` before running locally: the config defaults (`postgres` user / `postgres` db) do **not** match what `docker-compose.yml` provisions (`joaquin` / `hospital-db`).
- The build binary is `bin/hotel-api` in the Makefile (copy-paste leftover from a hotel project) — leave it.
- Migrations auto-run at startup via `database.AutoMigrate` (`internal/shared/infrastructure/database/database.go`); there is no migration tool.

### Backend gotchas

- All routes live under `/api`. `/health` is public. Everything else behind JWTAuth; patients, specialties, and doctors additionally require role `admin` (`cmd/api/main.go:99-108`).
- Flow: register a user with `"role": "admin"` → login for a JWT → send `Authorization: Bearer <token>`. `requests.http` is the quickest way to exercise endpoints.
- Code comments and API error messages are in **Spanish** — keep new ones in Spanish.
- Bruno env (`backend/Hospital/environments/hospitalEnv.yml`) embeds a hardcoded, expiring JWT; don't treat it as a fixture.

## Frontend (run from `frontend/`)

- `npm run dev` — dev server on `:3000`.
- `npm run sass` / `npm run sass-min` — compile `public/assets/scss/` → `public/assets/css/`.
- **Do not use `npm run build` on Linux/macOS** — it pipes through Windows `findstr` (`package.json:7`), which doesn't exist. Use `npx next build`.
- `next.config.ts` sets `output: "export"`, `trailingSlash: true`, a prod `basePath`, and ignores TS/ESLint errors at build time (so `next build` will not surface type errors).
- Both `package-lock.json` and `pnpm-lock.yaml` are committed — don't mix package managers.
- ESLint rules: double quotes, semicolons, tabs, camelCase.
- `@/*` path alias maps to the frontend root (`tsconfig.json`).
- Sidebar menu data lives in `shared/layouts-components/sidebar/nav.tsx` (`MENUITEMS`); it already lists `Doctores`/`Pacientes`/`Usuarios` pointing to `/components/doctores`, `/components/pacientes`, `/components/usuarios`. Active state is matched by URL path, so a page with `Contentlayout` lights up automatically.

### Frontend gotchas

- Pages under `frontend/pages/**` set `Component.layout = "Contentlayout"` (resolved dynamically in `_app.tsx` from the `layouts` map). Bare pages (e.g. `index.tsx`) render no layout — don't forget `PrelineScript` on those.
- Next.js pages router allows `pages/components/doctores.tsx` (route `/components/doctores`) to coexist with `pages/components/doctores/nuevo.tsx` (route `/components/doctores/nuevo`) — file acts as the segment index.
- Doctor feature pages (mock data only, not wired to the Go API):
  - `pages/components/doctores.tsx` — list table (Nombre, Apellido, DNI, Correo, N° Licencia) rendered with `Spktables`; has a "Nuevo Doctor" button linking to `/components/doctores/nuevo`.
  - `pages/components/doctores/nuevo.tsx` — create-doctor form.
  - `shared/data/doctorsdata.tsx` — mock `Doctor` array (id, firstName, lastName, dni, email, licenseNumber). Backend DTO field names differ: the API returns `first_name`, `last_name`, `email`, `dni`, `role`, `speciality_id`, `license_number`.
- No form library: forms use `useState` + a `changeHandler` (`e.target.name`), inputs with class `form-control`, labels `form-label text-defaulttextcolor`, required marker `<sup className="text-xs text-danger">*</sup>`. There are no reusable form components — write markup inline.
- `Spktables` (`shared/@spk-reusable-components/tables/spk-tables.tsx`) is the table component: `header` is an array of `{ title, headerClassname? }`; rows are `children`. `spk-gridjstable.tsx` (GridJS) is broken — `gridjs-react` isn't installed.