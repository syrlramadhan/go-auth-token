(function () {
	'use strict'

	var forms = document.querySelectorAll('.needs-validation')

	Array.prototype.slice.call(forms)
		.forEach(function (form) {
			form.addEventListener('submit', function (event) {
				if (!form.checkValidity()) {
					event.preventDefault()
					event.stopPropagation()
				}

				form.classList.add('was-validated')
			}, false)
		})
})()

async function createUser(name, email, password) {
	const url = 'http://localhost:8080/api/user/create';
	const data = {
		name: name,
		email: email,
		password: password
	};

	try {
		const response = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(data)
		});

		const result = await response.json();

		if (response.ok) {
			console.log('User created successfully:', result);

		} else {
			console.error('Error creating user:', result);
		}

		window.location.href = 'index.html';
	} catch (error) {
		console.error('Error:', error);
	}
}

document.getElementById('register-form').addEventListener('submit', function (event) {
    event.preventDefault();
    const form = this;
    if (!form.checkValidity()) {
        event.stopPropagation();
    } else {
        const name = document.getElementById('name').value;
        const email = document.getElementById('register-email').value;
        const password = document.getElementById('register-password').value;
        createUser(name, email, password);
    }

    form.classList.add('was-validated');
});