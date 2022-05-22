$(function() {
    var code, brand = "", pfc,
        serie, size, model, price, location, note;
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

    $('input[type="radio"]input[checked]').parent().css("background-color", "#3ebf4c")
    $('input[type="radio"]input[checked]').parent().css("border", "solid")

    function setCode() {
        setVariables();

        numb = code.substr(code.length - 2, code.length);
        if (code.indexOf("-") == -1) {
            code = code.substr(0,8)+ size+code.substr(code.length - 3, code.length);
            $('#code').val(code);
            return;
        } 
        
        code = code.split("-")[0] + "-";

        if (brand === undefined) code += "";
        else {
            if (brand.length >= 3) {
                code += brand.substr(0, 3)
            } else {
                code += brand;
            }
        }

        if (size === undefined) code += "";
        else code += size;


        if (location === undefined || location == "none") code += "";
        else {
            code += location.substr(0, 3);
        };
        console.log(code)
        code += numb;
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
        var size = Number($('#size').val());
        if (isNaN(Number(size))) {
            alert("La talla debe ser un numero");
            $('#size').val("");
            setCode();
        }
    });

    $('#model').on('input', function() {
        $('#model').val($('#model').val().toUpperCase());
        // setCode();
    });

    $('#price').on('input', setCode);
    $('#price').on('change', function() {
        var price = Number($('#price').val());
        if (price <= 0) {
            alert("El precio no debe ser menor a 0");
            $('#price').val("");
        }
    });

    $('input[type=radio][name=location]').change(function() {
        setCode();
    });



    function setVariables() {
        code = $('#code').val();
        brand = $('#brand').val();
        serie = $('#serie').val();
        size = $('#size').val();
        model = $('#model').val();
        price = $('#price').val();
        location = $('input[name="location"]:checked').val()
        note = $('#note').val();
        pfc = $('#pfc').val();

        if (pfc === "") {
            pfc = "S/N";
        }
        return {
            code: code,
            pfc: pfc, 
            brand: brand,
            serie: serie,
            size: size,
            model: model,
            price: price,
            location: location
        }
    }


    $('input[type="radio"]').click(function(){
        $(".locOption").css("background-color", "#ecf0f5")
        $(".locOption").css("border", "none")
        if ($(this).is(':checked'))
        {
        $('input[type="radio"]input[value="'+$(this).val()+'"]').parent().css("background-color", "#3ebf4c")
        $('input[type="radio"]input[value="'+$(this).val()+'"]').parent().css("border", "solid")
        
        }
      });


    $("form").submit(function(event) {
        setVariables();
        if ($('#location').val() == "none") {
            alert("No se ha seleccionado ubicación");
            event.preventDefault();
            return
        }
        event.preventDefault();
        // $('#confirmModal').modal("show");
         var data = {
            code: code,
            pfc: pfc,
            brand: brand,
            serie: serie,
            model: model,
            price: Number(price),
            size: Number(size),
            location: location,
            note: note
        };
        var realcode = $('#realcode').text();
        $.ajax({
            type: "PUT",
            url: url_server + "/producto/" + realcode,
            data: data,
            success: function() {
                window.location.href = "/producto/" + code;
            },
            fail: function() {
                window.location.href = "/producto/" + realcode;
            },
            complete: function(response) {
                if(response.status == 403) window.location.reload()
                else  window.location.href = "/producto/" + code;
            }
        }).done(function() {
            window.location.href = "/producto/"+code
        });;
    });

    // $("#confirm").on('click', function(event) {
    //     event.preventDefault();
    //     setVariables();
    //     var password = $('#password').val();

    //     if (password === "") {
    //         alert("Debe ingresar la contraseña");
    //         return;
    //     }
       
    // });

});