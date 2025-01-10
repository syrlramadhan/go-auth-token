// Mendapatkan data user dari API menggunakan token
async function getUserData() {
    const token = localStorage.getItem('authToken');  // Ambil token dari localStorage

    if (!token) {
        console.error('User is not authenticated');
        alert('You need to log in first');
        window.location.href = 'index.html';  // Arahkan ke index.html jika belum login
        return;
    }

    const url = 'http://localhost:8080/api/user/me';

    try {
        const response = await fetch(url, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`,  // Menambahkan token di header untuk autentikasi
            }
        });

        const result = await response.json();
        if (response.ok) {
            console.log('User data:', result);
            document.getElementById('user-name').textContent = result.data.name;
            document.getElementById('user-email').textContent = result.data.email;
            window.userId = result.data.id;
        } else {
            console.error('Error fetching user data:', result);
        }
    } catch (error) {
        console.error('Error:', error);
    }
}

// Memuat data user setelah halaman siap
document.addEventListener('DOMContentLoaded', function () {
    const token = localStorage.getItem('authToken');
    if (!token) {
        alert('You must be logged in to view your profile');
        window.location.href = 'index.html';  // Arahkan ke halaman login jika belum login
    } else {
        getUserData();  // Ambil data pengguna jika sudah login
    }
});

// Fungsi untuk logout dan menghapus token dari localStorage
function logout() {
    localStorage.removeItem('authToken');  // Hapus token dari localStorage
    console.log('Logged out successfully');
    window.location.href = 'index.html';  // Arahkan ke halaman login setelah logout
}

// Fungsi untuk mengupdate data user
async function updateUserData(name, email) {
    const url = `http://localhost:8080/api/user/update/${window.userId}`;  // URL untuk update data user berdasarkan ID
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
            getUserData();  // Perbarui data setelah berhasil diupdate
        } else {
            console.error('Error updating user:', result);
        }
    } catch (error) {
        console.error('Error:', error);
    }
}

// Fungsi untuk menghapus user
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
            window.location.href = 'index.html';  // Arahkan ke halaman login setelah penghapusan
        } else {
            console.error('Error deleting user:', result);
        }
    } catch (error) {
        console.error('Error:', error);
    }
}

// Event listener untuk tombol update
document.getElementById('update-button').addEventListener('click', function () {
    const name = prompt('Enter new name', document.getElementById('user-name').textContent);
    const email = prompt('Enter new email', document.getElementById('user-email').textContent);

    if (name && email) {
        updateUserData(name, email);  // Kirim data untuk update
    }
});

// Event listener untuk tombol delete
document.getElementById('delete-button').addEventListener('click', function () {
    const confirmation = confirm('Are you sure you want to delete your account?');
    if (confirmation) {
        deleteUser();  // Hapus user
    }
});
