/**
 * 常用函数
 */
import fs from 'file-saver'
import XLSX from 'xlsx'


// 导出Excel表格
export function exportExcel (json, fields, filename) {
    json.forEach(item => {
        for (let i in item) {
            console.log(i, item)
            if (fields.hasOwnProperty(i)) {
                item[fields[i]] = item[i];
            }
            delete item[i]; //删除原先的对象属性
        }
    })

    filename = filename + '_' + new Date().getTime();

    let sheetName = filename //excel的文件名称
    let wb = XLSX.utils.book_new()  //工作簿对象包含一SheetNames数组，以及一个表对象映射表名称到表对象。XLSX.utils.book_new实用函数创建一个新的工作簿对象。
    let ws = XLSX.utils.json_to_sheet(json, { header: Object.values(fields) }) //将JS对象数组转换为工作表。
    wb.SheetNames.push(sheetName)
    wb.Sheets[sheetName] = ws
    const defaultCellStyle = { font: { name: "Verdana", sz: 13, color: "FF00FF88" }, fill: { fgColor: { rgb: "FFFFAA00" } } };//设置表格的样式
    let wopts = { bookType: 'xlsx', bookSST: false, type: 'binary', cellStyles: true, defaultCellStyle: defaultCellStyle, showGridLines: false }  //写入的样式
    let wbout = XLSX.write(wb, wopts)
    let blob = new Blob([s2ab(wbout)], { type: 'application/octet-stream' })
    fs.saveAs(blob, filename + '.xlsx')
}

const s2ab = s => {
    if (typeof ArrayBuffer !== 'undefined') {
        let buf = new ArrayBuffer(s.length)
        let view = new Uint8Array(buf)
        for (let i = 0; i != s.length; ++i) view[i] = s.charCodeAt(i) & 0xff
        return buf
    } else {
        let buf = new Array(s.length);
        for (let i = 0; i != s.length; ++i) buf[i] = s.charCodeAt(i) & 0xFF;
        return buf;
    }
}

export function DateIntToDate(value) {
    if (!value) return ''
    var date = value.toString();
    return date.slice(0,4) + '-' + date.slice(4,6) + '-' + date.slice(6,8);
}

export function parseUrl(query) {
    var isArray = function (obj) {
        return Object.prototype.toString.call(obj) === "[object Array]";
    };
    var url = Object.keys(query).map(function (key) {
      if (isArray(query[key])) {
        var iUrl = "";
        query[key].forEach((item) => {
          iUrl = iUrl + key + "[]=" + item + "&";
        });
        return iUrl.trim("&");
      } else {
        return (
          encodeURIComponent(key) + "=" + encodeURIComponent(query[key])
        );
      }
    })
    .join("&");

    return url;
}