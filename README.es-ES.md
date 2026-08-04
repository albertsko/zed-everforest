

# 🌲 Everforest para Zed

Esquema de colores [Everforest](https://github.com/sainnhe/everforest) para Zed.
El tema viene en las variantes `regular`, `material` y `blur`.

### Desarrollo

Los archivos en `assets/` son la fuente de verdad. `generate.go` combina las paletas y plantillas en los archivos confirmados bajo `themes/`. No edites los temas generados directamente.

El generador requiere Go.

| Comando         | Propósito                                                                       |
| --------------- | ------------------------------------------------------------------------------- |
| `make generate` | Regenerar cada archivo de tema.                                                 |
| `make ci`       | Ejecutar verificaciones de formato, vet y pruebas.                              |
| `make copygen`  | Regenerar los temas y copiarlos en `~/.config/zed/themes/` para pruebas locales.|

Al contribuir, actualiza `assets/`, ejecuta `make generate`, ejecuta `make ci` y realiza el commit tanto de los cambios en el código fuente como de los generados.

### Licencia

Publicado bajo la [Licencia MIT](LICENSE).
