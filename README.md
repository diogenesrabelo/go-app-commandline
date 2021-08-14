# go-app-commandline
Aplicação simples em go de linha de comando

Aplicação busca IPs e nomes dos servidores pelo host do site através de uma aplicação de linha de comando.

Exemplo para buscar os IPs:<br>
***go run main.go ip --host amazon.com.br*** <br>
Defaul busca o Ip do google.com:<br> ***go run main.go ip***

Exemplo para buscar os nomes dos servidores:<br>
***go run main.go servidor --host amazon.com.br*** <br>
Default busca o nome dos servidores do google.com:<br> ***go run main.go servidor***