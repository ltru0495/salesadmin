$(function() {

    var options = {
        seriesBarDistance: 15,
        height: '280px',  
    };


    $.getJSON(url_server + '/api/graficos/marcas', function(brands) {
        var labels = [];
        var series = [
            []
        ];

        for (var i = 0; i < brands.length; i++) {
            console.log(brands)
            labels.push(brands[i].name)
            series[0].push(brands[i].amount);
        }

        var data = {
            labels: labels,
            series: series,
           
            chartPadding: {
                top: 40,
                right: 0,
                bottom: 40,
                left: 0
            },
            axisY: {
                onlyInteger: true
            },
        }
        new Chartist.Bar('#cbrands', data, options);
        // new Chartist.Bar('#cbrands', data)
    });

    $.getJSON(url_server + '/api/graficos/series', function(series) {
        var labels = [];
        var seriesData = [
            []
        ];

        var aux;
        for (var k = 0; k < series.length; k++) {
            for (var i = k + 1; i < series.length; i++) {
                if (series[k].name >= series[i].name) {
                    aux = series[i];
                    series[i] = series[k];
                    series[k] = aux;
                }
            }
        }
        for (var i = 0; i < series.length; i++) {
            labels.push(series[i].name)
            seriesData[0].push(series[i].amount);
        }




        data = {
            labels: labels,
            series: seriesData,
            chartPadding: {
                top: 0,
                right: 0,
                bottom: 0,
                left: 0
            },
            axisY: {
                onlyInteger: true
            },
        }
        new Chartist.Bar('#cseries', data, options)
    });

})