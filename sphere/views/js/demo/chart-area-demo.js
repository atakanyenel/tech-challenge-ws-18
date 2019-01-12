// Set new default font family and font color to mimic Bootstrap's default styling
Chart.defaults.global.defaultFontFamily = '-apple-system,system-ui,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif';
Chart.defaults.global.defaultFontColor = '#292b2c';

// Area Chart Example

$.get("/usage-by-day", function (data, status) {

  var lineChartData = {
    labels: [],
    datasets: []
  };
  values = data["values"]
  lineChartData.labels=data["labels"]
  console.log(values)

  Object.keys(values).forEach(function (key, index) {
    newDataset = {
      label: "",
      borderColor: [],
      backgroundColor: [],
      fill: false,
      data: [],
      yAxisID: "y-axis-1"
    }
    var colorNames = Object.keys(window.chartColors);
    var colorName = colorNames[index % colorNames.length];
    var newColor = window.chartColors[colorName];
    newDataset.backgroundColor.push(newColor)
    newDataset.borderColor.push(newColor)
    console.log(key)
    newDataset.label = key
    Object.keys(values[key]).forEach(date => {
      console.log(values[key][date])
      newDataset.data.push(values[key][date])
    })
    lineChartData.datasets.push(newDataset)
  })
  var ctx = document.getElementById("myAreaChart");
  myLineChart = Chart.Line(ctx, {

    data: lineChartData,
    options: {
      responsive: true,
      hoverMode: 'index',
      stacked: false,
      title: {
        display: true,
        text: 'Daily Usage of Sockets'
      },
      scales: {
        yAxes: [{
          type: 'linear', // only linear but allow scale type registration. This allows extensions to exist solely for log scale for instance
          display: true,
          position: 'left',
          id: 'y-axis-1',
        }],
      }
    }
  });



})
