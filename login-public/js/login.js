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

async function loginUser(email, password) {
    const url = 'http://localhost:8080/api/user/login';
    const data = { email, password };

    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        });

        const result = await response.json();
        if (response.ok) {
            console.log('Login successful:', result);

            // Menyimpan token di localStorage
            localStorage.setItem('authToken', result.token); // Pastikan server mengembalikan token di field 'data.token'

            // Arahkan ke halaman user.html setelah login
            window.location.href = 'user.html';
        } else {
            console.error('Login failed: ', result);
            alert('Login failed: ' + result.message);
        }
    } catch (error) {
        console.error('Error:', error);
        alert('An error occurred during login. Please try again later.');
    }
}

async function getUserData() {
	const url = 'http://localhost:8080/api/user';

	try {
		const response = await fetch(url, {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json'
			}
		});

		const result = await response.json();
		if (response.ok) {
			console.log('User data:', result);
		} else {
			console.error('Error fetching user data:', result);
		}
	} catch (error) {
		console.error('Error:', error);
	}
}

async function updateUser(userId, name, email, password) {
	const url = `http://localhost:8080/api/user/update/${userId}`;
	const data = { name, email, password };

	try {
		const response = await fetch(url, {
			method: 'PUT',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(data)
		});

		const result = await response.json();
		if (response.ok) {
			console.log('User updated successfully:', result);
		} else {
			console.error('Error updating user:', result);
		}
	} catch (error) {
		console.error('Error:', error);
	}
}

async function deleteUser(userId) {
	const url = `http://localhost:8080/api/user/delete/${userId}`;

	try {
		const response = await fetch(url, {
			method: 'DELETE',
			headers: {
				'Content-Type': 'application/json'
			}
		});

		const result = await response.json();
		if (response.ok) {
			console.log('User deleted successfully:', result);
		} else {
			console.error('Error deleting user:', result);
		}
	} catch (error) {
		console.error('Error:', error);
	}
}

// Mengambil elemen form dan tombol login
const loginForm = document.getElementById('login-form');

// Event listener untuk menangani submit form
loginForm.addEventListener('submit', function (event) {
    event.preventDefault();
    const form = this;

    // Validasi form
    if (!form.checkValidity()) {
        event.stopPropagation();
    } else {
        const email = document.getElementById('login-email').value;
        const password = document.getElementById('login-password').value;
        loginUser(email, password);  // Panggil fungsi loginUser
    }

    form.classList.add('was-validated');
});