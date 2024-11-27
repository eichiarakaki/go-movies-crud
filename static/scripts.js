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

