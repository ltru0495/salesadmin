$(function() {
    $('.select2-single').select2();
    $("form")[0].reset();


    var code, brand = "",
        serie, size, model, price, quantity, location, note, pfc;
    var flag = true;

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

    function generateCode(size) {
        var code = "";

        var counter = $('#counter').val();
        code += counter + "MA";


        if (size === undefined) code += "";
        else code += size;



        return code;
    }

    $('input[type="radio"]').click(function(){
        $(".locOption").css("background-color", "#ecf0f5")
        $(".locOption").css("border", "none")
        if ($(this).is(':checked')) {
            $('input[type="radio"]input[value="'+$(this).val()+'"]').parent().css("background-color", "#3ebf4c")
            $('input[type="radio"]input[value="'+$(this).val()+'"]').parent().css("border", "solid")
        }
    });
      
    $('#brand').on('input', function() {
        $('#brand').val($('#brand').val().toUpperCase());
    });
    $('#brand').on('change', function() {
        $('#brand').val($('#brand').val().toUpperCase());
    });
    $('#serie').on('change', function() {
        var re = /^\d{2}\-\d{2}$/
        var ser = $('#serie').val();
        if (re.test(ser)) {
            var m = Number(ser.split('-')[0]);
            var M = Number(ser.split('-')[1]);
            if (isNaN(m) || isNaN(M)) {
                $('#sizes').empty();
                alert("La serie debe tener el formato: mm-MM");
                $('#serie').val("");
            } else {
                if (m >= M) {
                    $('#sizes').empty();
                    alert("El primer numero debe ser menor que el segundo");
                    $('#serie').val("");
                } else {
                    $('#sizes').empty();
                    appendSizes(m, M);
                }


            }
        } else {
            alert("La serie debe tener el formato: mm-MM");
            $('#serie').val("");
        }

    });

    $('#size').on('change', function() {
        var size = Number($('#size').val());
        if (size <= 20) {
            alert("La talla debe ser mayor que 20");
            $('#size').val("");
        }
    });

    $('#model').on('input', function() {
        $('#model').val($('#model').val().toUpperCase());
    });
    $('#model').on('change', function() {
        $('#model').val($('#model').val().toUpperCase());
    });
    $('#pfc').on('change', function() {
        $('#pfc').val($('#pfc').val().toUpperCase());
    });

    $('#price').on('change', function() {
        var price = Number($('#price').val());
        if (price <= 0) {
            alert("El precio no debe ser menor a 0");
            $('#price').val("");
        }
    });



    function setVariables() {
        brand = $('#brand').val();
        serie = $('#serie').val();
        size = $('#size').val();
        model = $('#model').val();
        price = $('#price').val();
        quantity = $('#quantity').val();
        location = $('input[name="location"]:checked').val()
        note = $('#note').val();
        pfc = $('#pfc').val();

        if (pfc === "") {
            pfc = "S/N";
        } 
    }

    function appendSizes(min, max) {
        for (var i = min; i <= max; i++) {
            $('#sizes').append(`
   				<div class="form-group row">
                    <div class="col-md-1"></div>
   					<div class="col-md-1">
   						<label class="size-label">` + i + `</label>
   					</div>
   					<div class="col-md-1">
   						<input type="text" class="size-input" placeholder="0" id="size_` + i + `">
   					</div>

   				</div>
   				`)
        }

        $('.size-input').on('change', function() {
            var q = Number(this.value);
            console.log(q)
            if (isNaN(q)) {
                alert("La cantidad debe ser un numero");
            } else if (q < 0) {
                alert("La cantidad debe ser mayor o igual a 0");
                this.value = 0;
            }
        });
    }

    //  generateCode(brand, serie, model, size, location

    $("form").submit(function(event) {
        setVariables();
        //HERE GOES VALIDATION
        if ($('input[name="location"]:checked').val() === undefined) {
            alert("No se ha seleccionado ubicación");
            event.preventDefault();
            return;
        }

        if (isNaN(Number($('#price').val()))) {
            alert("El precio debe ser un numero");
            event.preventDefault();
            return;
        }

        var ser = $('#serie').val();
        var m = Number(ser.split('-')[0]);
        var M = Number(ser.split('-')[1]);

        var products = [];
        var product = {};
        //Get all products
        //Global variables must be OK

        var size, quantity, code;
        for (var i = m; i <= M; i++) {
            size = i;
            quantity = $('#size_' + i).val();
            if (quantity == "") quantity = "0";
            code = generateCode( size);
            product = {
                code: code,
                brand: brand,
                serie: serie,
                model: model,
                price: Number(price),
                quantity: Number(quantity),
                size: Number(size),
                location: location,
                note: note,
                pfc: pfc
            };
            if (product.quantity != 0) {
                products.push(product);
            }
        }

        $.ajaxSettings.traditional = true;
        $.ajax({
            url: "/productos",
            method: "post",
            data: {
                data: JSON.stringify(products)
            },
            success: function() {
                window.location.href = "/producto/nuevo";
            },
            done: function() {
                window.location.href = "/producto/nuevo";
            },
            complete: function() {
                window.location.href = "/producto/nuevo";
            }
        });


        event.preventDefault();

    });

});