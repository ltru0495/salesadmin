$(function() {

    var action = "";
    $('#infotable').DataTable({
        "language": {
            "lengthMenu": "",
            "zeroRecords": "No se encontraron registros",
            "info": "",
            "infoEmpty": "No hay información disponible",
            "infoFiltered": "(Filtrado de _MAX_ en total)",
            "search": "Buscar",
            "paginate": {
                "first": "",
                "last": "",
                "next": "",
                "previous": ""
            },
        },
        "iDisplayLength": 10000
    });

    $('.paginate_button').css('display', 'none')


    $('.refund-button').click(function() {
        id = this.id;
        if (confirm("Seguro que desea registrar devolución la venta?")) {
            // $('#confirmModal').modal('show');
            // var password = $('#password').val();
            // if (password === "") {
            //     alert("Debe ingresar la contraseña");
            //     return;
            // }
            // var data = {
            //     password: password,
            // };

            $.ajax({
                type: "POST",
                url: url_server + "/venta/" + id + "/devolver",
                // data: data,
                success: function() {
                    window.location.reload();
                },
                fail: function() {
                    window.location.reload();
                },
                complete: function() {
                    window.location.reload();
                }
            }).done(function() {
                window.location.reload();
            });;
        }
    });

    // $("#confirm").on('click', function(event) {
    //     event.preventDefault();
    //     var password = $('#password').val();
    //     if (password === "") {
    //         alert("Debe ingresar la contraseña");
    //         return;
    //     }
    //     var data = {
    //         password: password,
    //     };

    //     $.ajax({
    //         type: "POST",
    //         url: url_server + "/venta/" + id + "/devolver",
    //         data: data,
    //         success: function() {
    //             window.location.reload();
    //         },
    //         fail: function() {
    //             window.location.reload();
    //         },
    //         complete: function() {
    //             window.location.reload();
    //         }
    //     }).done(function() {
    //         window.location.reload();
    //     });;
    // });



});