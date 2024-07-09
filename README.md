# Arquitetura do Backend do E-commerce

Este documento descreve a arquitetura do backend de um projeto de e-commerce, detalhando o sistema de autenticação OAuth2, os tipos de usuários e suas permissões, e os fluxos de dados.

## Autenticação OAuth2

### Fluxo de Autenticação
- **Pedido de Autenticação**: O cliente envia um pedido de autenticação.
- **Headers**: O pedido inclui o cabeçalho `Authorization: Basic Auth`.
- **Corpo do Pedido**: O corpo do pedido contém `login` e `senha` no formato `x-www-form-urlencoded`.
- **Servidor de Autenticação**: O servidor valida as credenciais e gera um JWT.

## Usuários e Permissões

### SELLER
- Pode postar produtos para vender.
- Pode acompanhar suas vendas.
- Pode deletar seus próprios produtos.

### CLIENT
- Pode adicionar produtos ao carrinho.
- Pode fazer compras.
- Pode acompanhar suas próprias compras.

### ADMIN
- Acesso a todos os endpoints.
- Pode deletar usuários e produtos.

## Fluxos de Dados

### Endpoints Públicos
- **Lista de produtos**: Disponível sem autenticação.
- **Produto por ID**: Disponível sem autenticação.

### Endpoints Privados
- **Autenticação Requerida**: A maioria das rotas requer autenticação e está conectada a usuários específicos.

## Banco de Dados

### Produtos
- Tabela que armazena informações sobre os produtos.

### Usuários
- Tabela que armazena informações sobre os usuários e suas permissões.

### Pedidos
- Tabela que armazena informações sobre as compras (orders) realizadas.

## Exemplo de Fluxo de Usuário

### SELLER
- Loga no sistema.
- Posta um produto.
- Acompanha vendas.

### CLIENT
- Loga no sistema.
- Adiciona produtos ao carrinho.
- Faz uma compra.
- Acompanha suas ordens.

### ADMIN
- Loga no sistema.
- Gerencia usuários e produtos.
- Deleta conforme necessário.

---

Este diagrama descreve a arquitetura do backend do sistema de e-commerce, com foco nas interações e permissões de diferentes tipos de usuários e no fluxo de autenticação.
