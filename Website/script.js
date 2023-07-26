function getNum() {
    const val = document.querySelector('input').value;

    // Call the Go API
    fetch('http://localhost:8080/getNums', {
        method: 'GET',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },})
        .then(response => response.json())
        .then(data => displayResults(data))
        .catch(error => console.error('Error:', error));

}

function displayResults(results) {
    // Get the element to write the results to
    const cont = document.getElementById('results');

    cont.innerHTML = '';
    const resultDiv = document.createElement('div');
    resultDiv.innerHTML = `<strong>Number of occurencers of word "${results.word}":</strong> ${results.number}`;
    cont.appendChild(resultDiv);
}