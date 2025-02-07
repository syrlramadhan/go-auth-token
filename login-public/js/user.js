async function getUserData() {
    const token = localStorage.getItem('authToken');

    if (!token) {
        alert('You need to log in first');
        window.location.href = 'index.html';
        return;
    }

    const url = 'http://localhost:8080/api/user/me';

    try {
        const response = await fetch(url, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`,
            }
        });

        const result = await response.json();
        if (response.ok) {
            console.log(result)
            document.getElementById('user-name').textContent = result.data.name;
            document.getElementById('user-email').textContent = result.data.email;
            if (result.data.photo == "No Data") {
                document.getElementById('profile-picture').src = "images/user-117.webp";
            } else {
                document.getElementById('profile-picture').src = `http://localhost:8080/api/user/uploads/${result.data.photo}`;
            }
            window.userId = result.data.id;
        } else {
            console.error('Error fetching user data:', result);
        }
    } catch (error) {
        console.error('Error:', error);
    }
}

document.addEventListener('DOMContentLoaded', function () {
    const token = localStorage.getItem('authToken');
    if (!token) {
        alert('You must be logged in to view your profile');
        window.location.href = 'index.html';
    } else {
        getUserData();
    }
});

document.getElementById('profile-picture').addEventListener('click', () => {
    document.getElementById('file-input').click();
});

document.getElementById('file-input').addEventListener('change', async () => {
    const formData = new FormData();
    const fileInput = document.getElementById('file-input').files[0];

    if (!fileInput) {
        alert('Please select a file to upload.');
        return;
    }

    formData.append('photo', fileInput);

    const url = `http://localhost:8080/api/user/update/photo/${window.userId}`;
    try {
        const response = await fetch(url, {
            method: 'PUT',
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('authToken')}`
            },
            body: formData,
        });

        const result = await response.json();
        if (response.ok) {
            alert('Photo updated successfully');
            getUserData();
            const reader = new FileReader();
            reader.onload = function (e) {
                document.getElementById('profile-picture').src = e.target.result;
            };
            reader.readAsDataURL(fileInput);
        } else {
            console.error('Error uploading photo:', result);
        }
    } catch (error) {
        console.error('Error:', error);
    }
});



function logout() {
    localStorage.removeItem('authToken');
    console.log('Logged out successfully');
    window.location.href = 'index.html';
}

async function updateUserData(name, email) {
    const url = `http://localhost:8080/api/user/update/${window.userId}`;
    const data = { name, email };

    try {
        const response = await fetch(url, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        });

        const result = await response.json();
        if (response.ok) {
            console.log('User updated successfully:', result);
            alert('User updated successfully');
            getUserData();
        } else {
            console.error('Error updating user:', result);
        }
    } catch (error) {
        console.error('Error:', error);
    }
}

async function deleteUser() {
    const url = `http://localhost:8080/api/user/delete/${window.userId}`;

    try {
        const response = await fetch(url, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json',
            },
        });

        const result = await response.json();
        if (response.ok) {
            console.log('User deleted successfully:', result);
            alert('User deleted successfully');
            window.location.href = 'index.html';
        } else {
            console.error('Error deleting user:', result);
        }
    } catch (error) {
        console.error('Error:', error);
    }
}

document.getElementById('update-button').addEventListener('click', function () {
    const name = prompt('Enter new name', document.getElementById('user-name').textContent);
    const email = prompt('Enter new email', document.getElementById('user-email').textContent);

    if (name && email) {
        updateUserData(name, email);
    }
});

document.getElementById('delete-button').addEventListener('click', function () {
    const confirmation = confirm('Are you sure you want to delete your account?');
    if (confirmation) {
        deleteUser();
    }
});
