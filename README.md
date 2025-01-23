# Golang API Rest

## Descrição Geral
Este projeto é uma REST API desenvolvida em Golang, que permite gerenciar tarefas (tasks) em um banco de dados. A API utiliza o Prisma como ORM para facilitar a manipulação do banco de dados e inclui funcionalidades como criação, listagem, atualização e exclusão de tasks. Cada task possui um status que indica se está concluída (`completed`).

## Funcionalidades
- Adicionar tasks ao banco de dados.
- Atualizar informações de uma task específica.
- Marcar tasks como concluídas ou não.
- Listar todas as tasks.
- Remover tasks do banco de dados.

## Tecnologias Utilizadas
- **Golang**: Linguagem principal para desenvolvimento da API.
- **Prisma ORM**: Ferramenta para modelagem e manipulação do banco de dados.
- **SQLite**: Banco de dados local utilizado (pode ser substituído com base na `DATABASE_URL`).

## Configuração do Ambiente
### Variáveis de Ambiente
Certifique-se de configurar as seguintes variáveis no arquivo `.env`:
- `DATABASE_URL`: URL de conexão com o banco de dados.

### Prisma Schema
O schema do Prisma é utilizado para gerar o banco de dados e os modelos necessários. Certifique-se de que ele esteja configurado corretamente no arquivo `schema.prisma`.

## Estrutura do Banco de Dados
O banco de dados utilizado é local (SQLite). Ele contém a seguinte tabela principal:

### `tasks`
| Campo         | Tipo      | Descrição                           |
|---------------|-----------|-------------------------------------|
| `id`          | `int`     | Identificador único da task.        |
| `title`       | `string`  | Título ou descrição da task.        |
| `completed`   | `boolean` | Indica se a task foi concluída.     |

## Rotas e Endpoints
### Listar todas as tasks
**GET** `/tasks`  
Retorna uma lista de todas as tasks cadastradas.

### Criar uma nova task
**POST** `/tasks`  
Cria uma nova task.  
Exemplo de corpo da requisição (JSON):  
```json
{
  "title": "Minha nova task",
  "completed": true,
}
