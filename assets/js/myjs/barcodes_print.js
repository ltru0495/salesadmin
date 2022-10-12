$(function() {

    function showModal(str) {
        $('#optionModal').modal('show');
    }

    $('#code').on('input', function() {
        let code = $("#code").val();
        if (code.length >= 10) {
            $('#get').trigger("click");
        }
    })

    $('#get').on('click', (event) => {
        let code = $("#code").val();
	code = code.toUpperCase();
        $.getJSON(url_server + '/api/producto/' + code, function(product) {
            var values = [product.code, product.brand, product.serie, product.size, product.model, product.sprice]
            appendProductToTable($("#productstable"), values)
            $("#code").val("");
        }).fail(function() {
            // alert("PRODUCTO NO ENCONTRADO");
        });
    });
    // 000370-TNN40CAJ01
    function appendProductToTable(table, values) {
        var content = "";
        for (i = 0; i < values.length; i++) {
            content += "<td>" + values[i] + "</td>"
        }
        table.append('<tr >' + content + '</tr>')
    }

    $("#print").on("click", function() {
        let rows = $("tr");
        let products = [];
        for (var i = 1; i < rows.length; i++) {
            console.log(rows[i])
            products.push({
                code: rows[i].children[0].innerHTML,
                brand: rows[i].children[1].innerHTML,
                serie: rows[i].children[2].innerHTML,
                size: Number(rows[i].children[3].innerHTML),
                model: rows[i].children[4].innerHTML,
                sprice: Number(rows[i].children[5].innerHTML)
            })
        }

        if (products.length == 0) return

        $.ajaxSettings.traditional = true;
        $.ajax({
            url: "/api/productos/barcodes",
            method: "post",
            data: {
                data: JSON.stringify(products)
            },
        }).done(function(res) {
            if (res.status == 200) {
                window.open(res.message, '_blank');
            }
        });

    });

    $("#print2").on("click", function() {
        let rows = $("tr");
        let products = [];
        for (var i = 1; i < rows.length; i++) {
            console.log(rows[i])
            products.push({
                code: rows[i].children[0].innerHTML,
                brand: rows[i].children[1].innerHTML,
                serie: rows[i].children[2].innerHTML,
                size: Number(rows[i].children[3].innerHTML),
                model: rows[i].children[4].innerHTML,
                sprice: Number(rows[i].children[5].innerHTML)
            })
        }

        if (products.length == 0) return

        $.ajaxSettings.traditional = true;
        $.ajax({
            url: "/api/productos/barcodes2",
            method: "post",
            data: {
                data: JSON.stringify(products)
            },
        }).done(function(res) {
            if (res.status == 200) {
                window.open(res.message, '_blank');
            }
        });

    });

});
