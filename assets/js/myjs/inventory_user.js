const table = document.getElementById("table__content")
const pageNum = document.getElementById("pageNum")
const pageCount = document.getElementById("pageCount")
const prev = document.getElementById("prev")
const next = document.getElementById("next")


function createRow(p) {
    console.log(p.code);
    const tr = document.createElement("tr")
    tr.innerHTML = `
    <td class="table__cell">${p.brand}</td>
    <td class="table__cell">${p.pfc}</td>
    <td class="table__cell">${p.size}</td>
    <td class="table__cell">${p.model}</td>
    <td class="table__cell">${p.location}</td>
    <td class="table__cell">${p.note}</td>
    `

    return tr
}


let rxs = {}
const filters = ["brand", "pfc","size", "model", "location", "note"]

let currentPage = 1
let total = Number.MAX_VALUE
pageNum.value = currentPage
prev.disabled = true


function validatePage(page){
    return Number.isInteger(page) && page >= 1 && page <= total
}

function getURLParams(){
    filters.forEach(filter => {  
        const f = document.getElementById(`${filter}Query`)

        rxs[filter] = f.value .toUpperCase()
    })
}
function fetchData(page) {
    getURLParams()
    fetch(url_server+`/api/inventory/${page}?`+new URLSearchParams(rxs))
    .then(r => r.json())
    .then(response => {
        table.innerHTML = ""
        response.products.forEach(p => {
            table.appendChild(createRow(p))
        })
        total = response.pageCount

        currentPage = page
        checkPages()
        checkButtons()
    })
}
function checkButtons(){
    if (currentPage  == 1) prev.disabled = true 
    else prev.disabled = false
    if (currentPage  == total) next.disabled = true 
    else next.disabled = false
}
function checkPages(){
    pageCount.value = total
    pageNum.value = currentPage
}
fetchData(currentPage)

pageNum.addEventListener("keypress", function(event) {
    if (event.key === "Enter") {
        let page = parseInt(pageNum.value)
        
        if (!validatePage(page)) return
        fetchData(page)
    }
})

prev.addEventListener('click', (e)=>{
    e.preventDefault()
    let page = parseInt(pageNum.value)
    page--
    if (!validatePage(page)) return
    pageNum.value = page
    fetchData(page)
    
})
next.addEventListener('click', (e)=>{
    e.preventDefault()
    let page = parseInt(pageNum.value)
    page++
    console.log(page);
    
    if (!validatePage(page)) return
    pageNum.value = page
    fetchData(page)
})


filters.forEach(filter => {
    rxs[filter] = ""
    document.getElementById(`${filter}Query`).addEventListener('input', (event) => {
        currentPage = 1
        fetchData(currentPage)
    })
})

document.getElementById("search").addEventListener('click', ()=>{
    const filters = document.getElementById('filterRow')
    filters.classList.toggle('hidden')
})
function showModal(str) {
    JsBarcode("#barcode", str);
    $('#productModal').modal('show');
    currentCode = str;
}
var realcode;

function deleteProduct(str) {
    realcode = str;
    if (confirm("Seguro que desea eliminar el producto: " + str + "?")) {
        $.ajax({
            type: "DELETE",
            url: url_server + "/producto/" + realcode,
        }).done(function(res) {

            if(res.status == 200) {
                $('#'+realcode).remove()
                alert('Producto eliminado correctamente');
            } else {
                alert('Ha ocurrido un error')
            }
        });;
    }
}


$('#pdf').on('click', function(event) {
    event.preventDefault();

    ;
    if ($('input[name="type"]:checked').val() != "almacen") {
        var base = "/barcodes/pdf?code=" + currentCode + "&n=1&size=s";
    } else {
        var base = "/barcodes/pdf?code=" + currentCode + "&n=1&size=n";
    }

    $.ajax({
        url: base,
        method: "get",
    }).done(function(res) {
        if (res.status == 200 ){
            window.open(res.message, '_blank');
        }
    });
    $('#pdflink').attr('href', base)
});