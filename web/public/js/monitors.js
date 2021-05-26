$(function () {
    $.getJSON("/data/monitors", function (alerts) {
        debugger
        var dataGrid = $("#gridContainer").dxDataGrid({
            dataSource: alerts[0].ParsedJson.monitors,
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
                    dataField: "id",
                    caption: "Id",
                },
                {
                    dataField: "last_triggered_ts",
                    caption: "Date",
                },
                {
                    dataField: "name",
                },
                {
                    dataField: "classification",
                },
                {
                    dataField: "query",
                },
                {
                    dataField: "status",
                },
                {
                    dataField: "type",
                },
            ]
        }).dxDataGrid('instance');
    });
});