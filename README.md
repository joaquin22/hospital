# Funciones Stateful y Stateless

Clasificación de los componentes React del frontend (`frontend/`). Un componente **stateful** maneja estado local (hooks como `useState`/`useReducer`), estado global (Redux) o contexto (Context API). Un componente **stateless** recibe todo por props y renderiza sin estado propio (presentacional).

> Nota: el backend (Go) es **stateless** — API REST con autenticación JWT, sin sesiones en servidor; el estado persistente vive solo en PostgreSQL.

## Stateful

1. `shared/i18n/LanguageContext.tsx` — `LanguageProvider`: estado global de idioma (`lang` vía `useState`), persistencia en `localStorage` y sincronización del atributo `lang` del documento con `useEffect`. Expone el hook `useLanguage()` para consumir el contexto.
2. `pages/_app.tsx` — `App`: estado `pageloading` (`useState`) expuesto a todo el árbol mediante `Initialload.Provider` (Context API).
3. `pages/index.tsx` — landing page con estado local de formulario/newsletter (`err`, `data`, `backgroundClass`).
4. `pages/components/doctores/nuevo.tsx` — `NuevoDoctor`: formulario controlado con estado local `data` y `errors` (`useState`); handlers `changeHandler`, `validate` y `handleSubmit`.

## Stateless

1. Sin sesiones en servidor: la auth es por JWT firmado con HMAC-SHA256 (jwt_token_generator.go). El middleware JWTAuth valida firma + expiración en cada request de forma local, sin consultar un store de sesiones (auth.go); solo deja user_id/user_role en el contexto de Gin.
2. Cada request es independiente y reproducible: estado persistente solo en Postgres. No hay caché ni estado de módulo mutable; lo único "con estado" es infra configurable e inmutable (JWTTokenGenerator con secret/expiryMinutes), pero es config, no estado de request.