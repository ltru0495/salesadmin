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

    $('#infotablel').DataTable({
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

    var id;
    $('.delete-button').click(function() {
        action = "delete"
        id = this.id;
        if (confirm("Seguro que desea eliminar la venta?") ){
            $('#confirmModal').modal('show');
        }
    });

    $('.refund-button').click(function() {
        action = "refund"
        id = this.id;
        if (confirm("Seguro que desea registrar devolución la venta?") ){
            $('#confirmModal').modal('show');
        }
    });

    $("#confirm").on('click', function(event) {
        event.preventDefault();
        var password = $('#password').val();
        if (password === "") {
            alert("Debe ingresar la contraseña");
            return;
        }
        var data = {
            password: password,
        };
        
        var url_action ="";
        if (action =="refund") url_action = "devolver" 
        if (action =="delete") url_action = "eliminar"
        $.ajax({
            type: "POST",
            url: url_server + "/venta/"+id+"/"+url_action,
            data: data,
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
    });



});