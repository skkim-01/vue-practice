import PageFooter from '@/components/core/pagefooter';
import ToolBar from '@/components/core/toolbar';
import Navigation from '@/components/core/navigationdrawer';

// EasyDataTable
import 'vue3-easy-data-table/dist/style.css';
import Vue3EasyDataTable from 'vue3-easy-data-table';

// vuetify
import { vuetify3 } from "@/plugins/vuetify3";


function setComponents(app) {
    app.component('page-footer', PageFooter);
    app.component('toolbar', ToolBar);
    app.component('navigation', Navigation);

    app.component('EasyDataTable', Vue3EasyDataTable);

    app.use(vuetify3);
}


export {
    setComponents
}