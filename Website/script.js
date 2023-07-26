function getNum() {
    const val = document.querySelector('input').value;

    fetch('http://localhost:8080/getNums', {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ "id": val })
    })
   .then(response => response.json())
   .then(data => displayResults(data))
}

function displayResults(results) {
    // Get the element to write the results to
    const cont = document.getElementById('results');

    cont.innerHTML = '';
    const resultDiv = document.createElement('div');
    resultDiv.innerHTML = `<strong>Number of occurencers of word "${results.word}":</strong> ${results.number}`;
    cont.appendChild(resultDiv);
}