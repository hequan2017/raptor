<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="区域">
          <el-input v-model="searchInfo.region" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="类型">
          <el-input v-model="searchInfo.type" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="品牌">
          <el-input v-model="searchInfo.brand" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="型号">
          <el-input v-model="searchInfo.model" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="系统">
          <el-input v-model="searchInfo.system" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="主机名">
          <el-input v-model="searchInfo.hostname" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="资产编号">
          <el-input v-model="searchInfo.sn" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="一级负责人">
          <el-input v-model="searchInfo.first" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="内网IP">
          <el-input v-model="searchInfo.intranet" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="外网IP">
          <el-input v-model="searchInfo.public" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="状态">
          <el-input v-model="searchInfo.status" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="text" @click="deleteVisible = false">取消</el-button>
                <el-button size="small" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
<!--        <el-table-column align="left" label="日期" width="180">-->
<!--            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>-->
<!--        </el-table-column>-->
        <el-table-column align="left" label="区域" prop="region" width="120" />
        <el-table-column align="left" label="类型" prop="type" width="120" />
<!--        <el-table-column align="left" label="品牌" prop="brand" width="120" />-->
        <el-table-column align="left" label="型号" prop="model" width="120" />
<!--        <el-table-column align="left" label="系统" prop="system" width="120" />-->
        <el-table-column align="left" label="主机名" prop="hostname" width="120" />
        <el-table-column align="left" label="资产编号" prop="sn" width="120" />
<!--        <el-table-column align="left" label="一级负责人" prop="first" width="120" />-->
        <el-table-column align="left" label="内网IP" prop="intranet" width="120" />
        <el-table-column align="left" label="外网IP" prop="public" width="120" />
        <el-table-column align="left" label="状态" prop="status" width="120" />
        <el-table-column align="left" label="CPU" prop="cpu" width="120" />
        <el-table-column align="left" label="内存" prop="memory" width="120" />
        <el-table-column align="left" label="硬盘" prop="disk" width="120" />
<!--        <el-table-column align="left" label="其他" prop="other" width="120" />-->
<!--        <el-table-column align="left" label="上线时间" prop="uptime" width="120" />-->
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateAssetFunc(scope.row)">变更</el-button>
            <el-button type="text" icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="区域:">
          <el-input v-model="formData.region" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="产品线" prop="products">
          <el-select v-model="formData.products" placeholder="请选择下拉选择产品线" multiple clearable  value-key="ID"
                     :style="{width: '100%'}">
            <el-option v-for="(item, index) in productsOptions" :key="item.ID" :label="item.name"
                       :value="item"  :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="类型:">
          <el-input v-model="formData.type" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="品牌:">
          <el-input v-model="formData.brand" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="型号:">
          <el-input v-model="formData.model" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="系统:">
          <el-input v-model="formData.system" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="主机名:">
          <el-input v-model="formData.hostname" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="资产编号:">
          <el-input v-model="formData.sn" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="一级负责人:">
          <el-input v-model="formData.first" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="内网IP:">
          <el-input v-model="formData.intranet" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="外网IP:">
          <el-input v-model="formData.public" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:">
          <el-input v-model="formData.status" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="CPU:">
          <el-input v-model="formData.cpu" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="内存:">
          <el-input v-model="formData.memory" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="硬盘:">
          <el-input v-model="formData.disk" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="其他:">
          <el-input v-model="formData.other" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="上线时间:">
          <el-date-picker v-model="formData.uptime" type="date" style="width:100%" placeholder="选择日期" clearable />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'Asset'
}
</script>

<script setup>
import {
  createAsset,
  deleteAsset,
  deleteAssetByIds,
  updateAsset,
  findAsset,
  getAssetList
} from '@/api/asset'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import {getProductList} from "@/api/product";

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        region: '',
        products:[],
        type: '',
        brand: '',
        model: '',
        system: '',
        hostname: '',
        sn: '',
        first: '',
        intranet: '',
        public: '',
        status: '',
        cpu: '',
        memory: '',
        disk: '',
        other: '',
        uptime: new Date(),
        })

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const productsOptions = ref([])

const getProductsOptions = async() => {
  const table = await getProductList({ page: page.value, pageSize: 1000})
  if (table.code === 0) {
    table.data.list.forEach(function(val, index, arr){
      productsOptions.value.push(val)
    });
  }
}


// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getAssetList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteAssetFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteAssetByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateAssetFunc = async(row) => {
    getProductsOptions()
    const res = await findAsset({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reasset
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteAssetFunc = async (row) => {
    const res = await deleteAsset({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    getProductsOptions()
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    productsOptions.value = []
    formData.value = {
        region: '',
        type: '',
        products:[],
        brand: '',
        model: '',
        system: '',
        hostname: '',
        sn: '',
        first: '',
        intranet: '',
        public: '',
        status: '',
        cpu: '',
        memory: '',
        disk: '',
        other: '',
        uptime: new Date(),
        }
}
// 弹窗确定
const enterDialog = async () => {
      let res
      switch (type.value) {
        case 'create':
          res = await createAsset(formData.value)
          break
        case 'update':
          res = await updateAsset(formData.value)
          break
        default:
          res = await createAsset(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
        closeDialog()
        getTableData()
      }
}
</script>

<style>
</style>
