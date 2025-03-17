<template>
  <div class="app-container home">
    <el-row :gutter="20">
      <el-col :sm="24" :lg="24">
        <hr />
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :sm="24" :lg="12" class="left-column">
        <h2>{{ appTitle }}</h2>
        <p></p>
        <p>
          <b>当前版本:</b> <span>v{{ version }}</span>
        </p>
        <p>
          <el-button icon="HomeFilled" plain @click="goTarget('#')">访问主页</el-button>
        </p>
      </el-col>

      <el-col :sm="24" :lg="12" class="right-column">
        <el-row>
          <el-col :span="12">
            <h2>技术选型</h2>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="6">
            <h4>后端技术</h4>
            <ul>
              <li>golang</li>
              <li>goframe v2</li>
              <li>mysql</li>
              <li>golang-migrate</li>
              <li>...</li>
            </ul>
          </el-col>
          <el-col :span="6">
            <h4>前端技术</h4>
            <ul>
              <li>ruoyi-vue3</li>
              <li>element-plus</li>
              <li>xgplayer</li>
              <li>mpegts.js</li>
              <li>howler</li>
              <li>...</li>
            </ul>
          </el-col>
        </el-row>
      </el-col>
    </el-row>
    <el-divider />
  </div>
</template>

<script setup name="Index">
const appTitle = import.meta.env.VITE_APP_TITLE;
const version = import.meta.env.VITE_APP_VERSION;

function goTarget(url) {
  window.open(url, "__blank");
}
</script>

<style scoped lang="scss">
.left-column {
  padding-left: 20px;
}

.right-column {
  padding-left: 50px;
}

$font-family: "open sans", "Helvetica Neue", Helvetica, Arial, sans-serif;
$font-color: #676a6c;
$blockquote-border-color: #eee;

.home {
  font-family: $font-family;
  font-size: 13px;
  color: $font-color;
  overflow-x: hidden;

  blockquote {
    padding: 10px 20px;
    margin: 0 0 20px;
    font-size: 17.5px;
    border-left: 5px solid $blockquote-border-color;
  }

  hr {
    margin: 20px 0;
    border: 0;
    border-top: 1px solid $blockquote-border-color;
  }

  .col-item {
    margin-bottom: 20px;
  }

  ul {
    padding: 0;
    margin: 0;
    list-style-type: none;
  }

  h4 {
    margin-top: 0;
  }

  h2 {
    margin-top: 10px;
    font-size: 26px;
    font-weight: 100;
  }

  p {
    margin-top: 10px;

    b {
      font-weight: 700;
    }
  }

  .update-log {
    ol {
      display: block;
      list-style-type: decimal;
      margin: 1em 0;
      padding-inline-start: 40px;
    }
  }
}
</style>
