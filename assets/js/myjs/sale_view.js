$(function() {
    var seller, comment, price;


    function setVariables() {
        seller = $('#seller').val();
        comment = $('#comment').val();
        price= $('#price').val();
    }

    $('#seller').on('input', function (){
        $('#seller').val($('#seller').val().toUpperCase());
    });


    $("form").submit(function(event) {
        setVariables();
        if ( seller == "" || price == "" || isNaN(price)){
            alert("Ha ocurrido un error");
            event.preventDefault();
            return
        }
        event.preventDefault();
        $('#confirmModal').modal("show");
    });

    $("#confirm").on('click', function(event) {
        event.preventDefault();
        setVariables();
        var password = $('#password').val();
        var realid = $('#realid').text();

        if (password === "") {
            alert("Debe ingresar la contraseña");
            return;
        }
        var data = {
            password: password,
            price: Number(price),
            seller: seller,
            comment: comment,
        };

        $.ajax({
            type: "PUT",
            url: url_server + "/venta/" + realid,
            data: data,
            success: function() {
                window.location.href = "/venta/" + realid;
            },
            fail: function() {
                window.location.href = "/venta/" + realid;
            },
            complete: function(response) {
                window.location.href = "/venta/" + realid;
            }
        }).done(function() {
            window.location.href = "/inventario"
        });;
    });

});