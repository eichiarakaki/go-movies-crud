"use strict";

document.addEventListener('DOMContentLoaded', function () {
    console.log("DOM fully loaded and parsed"); // depuration
    fetch('/movies')
        .then(response => response.json())
        .then(data => {
            const moviesDiv = document.getElementById('movies');
            data.forEach(movie => {
                const movieDiv = document.createElement('div');
                movieDiv.classList.add('movie');
                movieDiv.innerHTML = `<h2>${movie.title}</h2><p>Directed by: ${movie.director.firstname} ${movie.director.lastname}</p><p>ISBN: ${movie.isbn}</p>`;
                moviesDiv.appendChild(movieDiv);
            });
        })
        .catch(error => console.error('Error fetching movies:', error));
});

document.getElementById('create-movie-form').addEventListener('submit', (e) => {
    e.preventDefault();
    const title = document.getElementById('create-title').value;
    const isbn = document.getElementById('create-isbn').value;
    const firstname = document.getElementById('create-firstname').value;
    const lastname = document.getElementById('create-lastname').value;

    // creating the JSON to request the data from the backend
    fetch('/movies', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            title: title,
            isbn: isbn,
            director: {
                firstname: firstname,
                lastname: lastname
            }
        })
    }).then(response => response.json())
        .then(data => {
            console.log('Movie Created:', data);
        })
});

document.getElementById('update-movie-form').addEventListener('submit', (e) => {
    e.preventDefault();
    const id = document.getElementById('update-id').value;
    const title = document.getElementById('update-title').value;
    const isbn = document.getElementById('update-isbn').value;
    const firstname = document.getElementById('update-firstname').value;
    const lastname = document.getElementById('update-lastname').value;

    // creating the JSON to request the data from the backend
    fetch(`/movies/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            title: title,
            isbn: isbn,
            director: {
                firstname: firstname,
                lastname: lastname
            }
        })
    }).then(response => response.json())
        .then(data => {
            console.log('Movie Updated:', data);
        }).catch(error => console.error('Error updating movie:', error));
});