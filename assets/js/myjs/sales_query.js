$(function() {
    $("#query").on('click', (e) => {
        e.preventDefault()
        let start = $("#dateStart").val()
        let end = $("#dateEnd").val();
        let location = $("#location").val()
        
        window.open(`${url_server}/api/ventas/exportar/${location}/${start}/${end}`)

    })
    // $("#query").on("click", function(e) {
    //     e.preventDefault();

    //     let range = $("#datefilter").val();

    //     if (range == "") {
    //         alert("No se ha seleccionado fecha")
    //         return;
    //     }


    //     console.log(url_server + `/api/ventas/${location}/${start}/${end}`);
    //     $.getJSON(url_server + `/api/ventas/${location}/${start}/${end}`, function(sales) {
    //         $('#infotable').DataTable().clear().destroy()

    //         $("#sales_body").empty()        
    //         sales.forEach(function(sale) {
    //             /*    <th>Hora</th>
    //                           <th>Marca</th>
    //                           <th>Modelo</th>
    //                           <th>Talla</th>
    //                           <th>Pertenece</th>
    //                           <th>P. Venta</th>
    //                           <th>P. Costo</th>
    //                           <th>Ganancia</th>
    //                           <th>Vendedor</th>
    //                           <th>Lugar de Venta</th>
    //                           <th>Comentario</th>*/
    //             $("#sales_body").append(`<tr>
    //           <td>${(new Date(sale.timestamp)).toLocaleTimeString()}</td>
    //           <td>${sale.brand}</td>
    //           <td>${sale.model}</td>
    //           <td>${sale.size}</td>
    //           <td>${sale.location}</td>
    //           <td>${sale.price}</td>
    //           <td>${sale.pricebuy}</td>
    //           <td>${sale.earning}</td>
    //           <td>${sale.seller}</td>
    //           <td>${sale.place}</td>
    //           <td>${sale.comment}</td>
    //           </tr>`);
    //         })


    //         $('#infotable').DataTable({
    //             "language": {
    //                 "lengthMenu": "",
    //                 "zeroRecords": "No se encontraron registros",
    //                 "info": "",
    //                 "infoFiltered": "(Filtrado de _MAX_ en )",
    //                 "search": "Buscar",
    //                 "paginate": {
    //                     "first": "",
    //                     "last": "",
    //                     "next": "",
    //                     "previous": ""
    //                 },
    //             },
    //             "iDisplayLength": 10000
    //         });
    //     });


    // })

});