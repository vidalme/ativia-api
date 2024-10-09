function onUpdate(username) {
    // Collect form data
    const formData = {
        id: document.getElementById('id').value,
        userName: document.getElementById('user_name').value,
        firstName: document.getElementById('first_name').value,
        lastName: document.getElementById('last_name').value,
        email: document.getElementById('email').value,
        password: document.getElementById('password').value,
        phone: document.getElementById('phone').value,
        userStatus: document.getElementById('user_status').value
    };

    // Perform a PUT request using fetch
    fetch('/user/'+username, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData) // Send the form data as JSON
    })
    .then(response => {
        if (response.ok) {
            alert('Usuário atualizado com sucesso!');
            window.location.href = '/'; // Redirect after success
        } else {
            alert('Erro ao atualizar usuário.');
        }
    })
    .catch(error => {
        console.error('Error:', error);
        alert('Ocorreu um erro ao tentar atualizar o usuário.');
    });
}