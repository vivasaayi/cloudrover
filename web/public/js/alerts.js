$(function () {
    $.getJSON("/data/alerts", function (alerts) {
        var dataGrid = $("#gridContainer").dxDataGrid({
            dataSource: alerts,
            columnsAutoWidth: true,
            showBorders: true,
            filterRow: {
                visible: true,
                applyFilter: "auto"
            },
            searchPanel: {
                visible: true,
                width: 240,
                placeholder: "Search..."
            },
            headerFilter: {
                visible: true
            },
            columns: [
                {
                    dataField: "Id",
                    caption: "Id",
                    width: 140
                },
                {
                    dataField: "Date",
                    alignment: "right",
                    width: 120,

                },
            ]
        }).dxDataGrid('instance');
    });
});