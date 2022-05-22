$(function() {
    var code, brand = "",
        serie, size, model, price, quantity, location, note;
    var flag = true;

    var acBrands = [];
    var acModels = [];


    var optionBrands = {
        url: "/api/marcas",
        getValue: "name",
        list: {
            match: {
                enabled: true
            },
            onSelectItemEvent: function() {
                $("#brand").trigger("change");
            },
            onHideListEvent: function() {
                $("#brand").trigger("change");
            }
        },
        theme: "square"
    };

    var optionModels = {
        url: "/api/modelos",
        getValue: "name",
        list: {
            match: {
                enabled: true
            },
            onSelectItemEvent: function() {
                $("#model").trigger("change");
            },
        },

        theme: "square"
    };

    $("#brand").easyAutocomplete(optionBrands);
    $("#model").easyAutocomplete(optionModels);

    function setCode() {
        setVariables();
        code = "";
        if (brand === undefined) code += "";
        else {
            if (brand.length >= 3) {
                code += brand.substr(0, 3)
            } else {
                code += brand;
            }
        }

        if (serie === undefined) code += "";
        else {
            code += serie.replace("-", "");
        }

        if (model === undefined) code += "";
        else {
            var aux = model.split(" ");
            code += aux[0];
        };

        if (size === undefined) code += "";
        else code += size;



        if (location === undefined || location == "none") code += "";
        else {
            code += location.substr(0, 3);
        };
        $('#code').val(code);
    }

    $('#brand').on('input', function() {
        $('#brand').val($('#brand').val().toUpperCase());
        setCode();
    });
    $('#brand').on('change', function() {
        $('#brand').val($('#brand').val().toUpperCase());
        setCode();
    });
    $('#serie').on('input', setCode);
    $('#serie').on('change', function() {
        var re = /^\d{2}\-\d{2}$/
        var ser = $('#serie').val();
        if (re.test(ser)) {
            var m = Number(ser.split('-')[0]);
            var M = Number(ser.split('-')[1]);;
            if (isNaN(m) || isNaN(M)) {
                alert("La serie debe tener el formato: mm-MM");
                $('#serie').val("");
            } else {

            }
        } else {
            alert("La serie debe ser del formato: mmMM");
            $('#serie').val("");
            setCode();
        }
    });

    $('#size').on('input', setCode);
    $('#size').on('change', function() {
        var size = Number($('#serie').val());
        if (isNaN(Number(size))) {
            alert("La talla debe ser un numero");
            $('#size').val("");
            setCode();
        }
    });

    $('#model').on('input', function() {
        $('#model').val($('#model').val().toUpperCase());

        setCode();
    });

    $('#price').on('input', setCode);
    $('#price').on('change', function() {
        var price = Number($('#price').val());
        if (price <= 0) {
            alert("El precio no debe ser menor a 0");
            $('#price').val("");
            setCode();
        }
    });

    $('#quantity').on('input', setCode);
    $('#location').on('input', setCode);



    function setVariables() {
        code = $('#code').val();
        brand = $('#brand').val();
        serie = $('#serie').val();
        size = $('#size').val();
        model = $('#model').val();
        price = $('#price').val();
        quantity = $('#quantity').val();
        location = $('#location').val();
        note = $('#note').val();
        return {
            code: code,
            brand: brand,
            serie: serie,
            size: size,
            model: model,
            price: price,
            quantity: quantity,
            location: location
        }
    }

    function getProduct() {
        brand = $('#brand').val();
        serie = $('#serie').val();
        size = $('#size').val();
        model = $('#model').val();
        price = $('#price').val();
        quantity = $('#quantity').val();
        location = $('#location').val();
    }


    $("form").submit(function(event) {
        if ($('#location').val() == "none") {
            alert("No se ha seleccionado ubicación");
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

        if (password==="") {
            alert("Debe ingresar la contraseña");
            return;
        }
        var data = {
            password: password, code : code, brand: brand, 
            serie: serie, model: model, price: Number(price),
            quantity: Number(quantity), size :Number(size), location:location, note: note
        };
        var realcode = $('#realcode').text();

        $.ajax({
            type: "PUT",
            url:  url_server+"/producto/"+realcode,
            data: data,
            success: function() {
                window.location.href ="/inventario"
            },
            fail: function(){
                window.location.href = "/inventario";  
            }, 
            stop: function() {
               window.location.href = "/inventario";   
            }
        }).done(function() {
            window.location.href ="/inventario"
        });;
    });

});
