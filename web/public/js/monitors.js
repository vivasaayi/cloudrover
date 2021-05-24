$(function () {
    $.getJSON("/data/monitors", function (alerts) {
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
                },
                {
                    dataField: "DateHappened",
                    caption: "Date",
                },
                {
                    dataField: "DeviceName",
                },
                {
                    dataField: "Host",
                },
                {
                    dataField: "AlertType",
                },
                {
                    dataField: "Payload",
                },
                {
                    dataField: "Priority",
                },
                {
                    dataField: "SourceTypeName",
                },
                {
                    dataField: "Text",
                },
                {
                    dataField: "Title",
                },
                {
                    dataField: "URL",
                },
            ]
        }).dxDataGrid('instance');
    });
});