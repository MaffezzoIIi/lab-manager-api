# Estratégia de Versionamento

Este projeto utiliza o [SemVer](https://semver.org/lang/pt-BR/) (Versionamento Semântico) para controle de versões.

## Formato das Versões

- **MAJOR**: Mudanças incompatíveis na API.
- **MINOR**: Novas funcionalidades compatíveis.
- **PATCH**: Correções de bugs e melhorias pequenas.

## Exemplos

- `v1.0.0`: Primeira versão estável.
- `v1.1.0`: Nova funcionalidade compatível.
- `v1.1.1`: Correção de bug.

## Quando criar Tags

| Situação              | Exemplo de Tag         |
|-----------------------|-----------------------|
| Primeira versão       | `v1.0.0`              |
| Correção de bugs      | `v1.0.1`, `v1.0.2`    |
| Nova funcionalidade   | `v1.1.0`, `v1.2.0`    |
| Mudança incompatível  | `v2.0.0`              |

## Como criar e subir uma tag

Execute os comandos abaixo no terminal:

```bash
git tag -a v1.0.0 -m "Primeira versão estável"
git push origin v1.0.0
