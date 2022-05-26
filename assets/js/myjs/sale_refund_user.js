$(function() {
    var id = "";
    $("form")[0].reset();

    $('#productButton').on('click', (event) => {
        event.preventDefault();
        var code = $("#code").val();
        // console.log(code)

        if (code === "") {
            alert("Codigo no especificado")
        } else {
            $(':input', '#saleform').not(':button, :submit, :reset, :hidden, #place').val('')
            $.getJSON(url_server + '/api/venta/' + code, function(sale) {
    			console.log(sale);

                id = sale._id;
                $('#code').val(sale.code);
                $('#brand').val(sale.brand);
                $('#serie').val(sale.serie);
                $('#size').val(sale.size);
                $('#place').val(sale.place);
                $('#location').val(sale.location);
                $('#model').val(sale.model);
                $('#price').val(sale.price);
                $('#seller').val(sale.seller);
                $('#comment').val(sale.comment);
                let rg =sale.timestamp.split("T")[0] 
// 000009MA21002
                console.log(rg);
                if ( rg != "" && rg != "01-01-0001"){
                    document.getElementById("date").value = rg

                }
            }).fail(function() {
                alert("VENTA NO ENCONTRADA");
            });

        }
    });

    function getField(fieldName) {
        return ($('#' + fieldName).val());
    }

    function setModalField(fieldName) {
        $('#modal-' + fieldName).text(getField(fieldName));
        if (fieldName === "price") $('#modal-' + fieldName).text("S/ " + getField(fieldName));
    }


    function getObjectToSend() {
        var sale = {
            code: getField('code'),
            seller: getField('seller'),
            price: getField('price'),
            place: getField('place'),
            location: getField('location'),
            comment: getField('comment')
        }
        return sale;
    }

    $('#save').on('click', (event) => {
        event.preventDefault();
        var sale = getObjectToSend();
        if (sale.code === "" || id == "") {
            alert("Venta no especificada")
            return;
        }
        if (confirm("¿Seguro que desea realizar una devolución de dinero?")) {
            $.ajax({
                type: "POST",
                url: url_server + "/venta/" + id + "/devolver",
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

    $('#change').on('click', (event) => {
        event.preventDefault();
        var sale = getObjectToSend();
        if (sale.code === "" || id == "") {
            alert("Venta no especificada")
            return;
        }
        if (confirm("¿Seguro que desea  realiza un cambio?")) {
            $.ajax({
                type: "POST",
                url: url_server + "/venta/" + id + "/cambiar",
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

    $("form").submit(function(event) {
        event.preventDefault();
        return;
    });
});