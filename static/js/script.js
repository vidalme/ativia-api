function onDelete(username) {
    let resposta = confirm("Tem certeza que deseja remover o usuário " + username + "?");
    if (resposta) {
        fetch("/user/" + username, {
            method: "DELETE",
            headers: {
                'Content-Type': 'application/json',
            },
        })
        .then(response => {
            if (response.ok) {
                alert("Usuário removido com sucesso!");
                window.location.reload();
            } else {
                alert("Erro ao remover o usuário.");
            }
        })
        .catch(error => {
            console.error("Error:", error);
            alert("Ocorreu um erro ao tentar remover o usuário.");
        });
    }
}