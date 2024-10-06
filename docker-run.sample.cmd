@REM copie para criar um container com postgres com o nome postgres-go e acompanhar seus logs (adicionar user, password e banco)
docker run --name postgres-go -e POSTGRES_USER= -e POSTGRES_PASSWORD= -e=POSTGRES_DB= -p 5432:5432 -d postgres:16.1 && docker logs --follow postgres-go
@REM o final do comando mantém o terminal observando os logs do servidor postgres

@REM copie para excluir o container postgres-go
docker rm -f postgres-go