#================================
#== DOCKER ENVIRONMENT
#================================
COMPOSE := @docker-compose

godev:
	${COMPOSE} -f docker-compose.local.yml up -d

#================================
#== GOLANG ENVIRONMENT
#================================
GO := @go
GIN := @gin

goinstall:
	${GO} get .