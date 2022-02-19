package controllers

import (
	"encoding/json"
	"errors"
	"flight/src/database"
	"flight/src/models"
	"flight/src/repository"
	"flight/src/response"
	"io/ioutil"
	"net/http"
)

// CriarAeroporto insere um aeroporto no banco de dados
func CriarAeronave(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	if !json.Valid([]byte(corpoRequest)) {
		response.Erro(w, http.StatusBadRequest, errors.New(" Registro enviado para servidor não compativel com o sistema. "))
		return
	}

	var aeronave models.Aeronave

	if erro = json.Unmarshal(corpoRequest, &aeronave); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = aeronave.Preparar("cadastro"); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	/* -- COLOCAR DUPLICIDADE DE AEROPORTO
	repositorioDuplicado := repository.NovoRepositoDeUsuarios(db)
	if codAnacLancto, _ := repositorioDuplicado.BuscarPor_CANAC_Email(usuario.Id_Anac, usuario.Email); codAnacLancto == "YES" {
		response.Erro(w, http.StatusBadRequest, errors.New(" Registro já informado para este EMAIL ou CANAC! Por favor verifique. "))
		return
	}*/

	repositorio := repository.NovoRepositoDeAeronave(db)
	aeronave.ID, erro = repositorio.Criar(aeronave)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, aeronave)
}
