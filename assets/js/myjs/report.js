$(function() {
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
    $('#code').val("");

    $('#code').on('input', function() {
        let code = $("#code").val();
        if (!isEmpty($("#" + code))) {
            $('#check').trigger("click");
        }
    })

    $('#check').on('click', (event) => {
        event.preventDefault();
        var code = $("#code").val();
        if (code != "") {
            // $("#cb_"+code).attr("checked", "checked");
            $("#" + code).css('color', '#fff');
            $("#" + code).css('background-color', 'rgb(0, 208, 84)');

            $.ajax({
                type: "POST",
                url: url_server + "/api/producto/reporte/" + code,
                success: function() {

                },
                fail: function() {},
                complete: function() {}
            }).done(function(res) {
                console.log(res)
            });;

            $("#code").val("");
        }
    });

    $('#validate').on('click', (event) => {
        event.preventDefault();

        if(!confirm("Desea validar la contabilidad del inventario?")){
            return;
        }

        window.open(url_server+"/productos/faltan", '_blank');
        $.ajax({
            type: "PUT",
            url: url_server + "/productos/reporte/validar",
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


    function isEmpty(obj) {
        for (var key in obj) {
            if (obj.hasOwnProperty(key))
                return false;
        }
        return true;
    }
});