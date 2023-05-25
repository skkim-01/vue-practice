// EasyDataTable
import 'vue3-easy-data-table/dist/style.css';
import Vue3EasyDataTable from 'vue3-easy-data-table';
import PageFooter from '@/components/core/pagefooter';
import ToolBar from '@/components/core/toolbar';
import Navigation from '@/components/core/navigationdrawer';

function setComponents(app) {
    app.component('EasyDataTable', Vue3EasyDataTable);
    app.component('page-footer', PageFooter);
    app.component('toolbar', ToolBar);
    app.component('navigation', Navigation);
}


export {
    setComponents
}