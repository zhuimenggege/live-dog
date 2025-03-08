<template>
  <div class="audio-player">
    <div class="controls">
      <button @click="togglePlay" class="play-button">
        <el-icon v-if="isPlaying">
          <VideoPause />
        </el-icon>
        <el-icon v-else>
          <VideoPlay />
        </el-icon>
      </button>

      <div class="time-progress-container">
        <div class="time current-time">{{ formatTime(currentTime) }}</div>
        <div class="progress-container" @click="seek">
          <div class="progress-bar" :style="{ width: progress + '%' }">
            <div class="progress-handle"></div>
          </div>
          <div class="progress-bg"></div>
        </div>
        <div class="time total-time">{{ formatTime(duration) }}</div>
      </div>

      <div class="volume-container">
        <el-icon>
          <Microphone />
        </el-icon>
        <div class="volume-slider-container">
          <input type="range" min="0" max="1" step="0.01" v-model="volume" @input="updateVolume"
            class="volume-slider" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount, watch } from 'vue';
import { Howl } from 'howler';
import { getToken } from '@/utils/auth';
import { VideoPlay, VideoPause, Microphone } from '@element-plus/icons-vue';

const props = defineProps({
  src: {
    type: String,
    required: true
  },
  autoplay: {
    type: Boolean,
    default: false
  }
});

const isPlaying = ref(false);
const progress = ref(0);
const currentTime = ref(0);
const duration = ref(0);
const volume = ref(1);

let sound = null;

const initializeAudio = () => {
  if (sound) {
    sound.unload();
  }

  sound = new Howl({
    src: [props.src],
    html5: false, // 修改为强制使用 Web Audio API
    autoplay: props.autoplay,
    volume: volume.value,
    format: ['mp3', 'aac'],
    xhr: {
      method: 'GET',
      headers: {
        Authorization: `Bearer ${getToken()}`,
      },
      withCredentials: true
    },
    onplay: handlePlay,
    onpause: handlePause,
    onend: handleEnd,
    onload: handleLoad,
    onloaderror: handleLoadError
  });
};

const handlePlay = () => {
  isPlaying.value = true;
  requestAnimationFrame(updateProgress);
};

const handlePause = () => {
  isPlaying.value = false;
};

const handleEnd = () => {
  isPlaying.value = false;
  progress.value = 100;
};

const handleLoad = () => {
  duration.value = sound.duration();
};

const handleLoadError = (id, error) => {
  console.error(`音频加载错误 (ID: ${id}):`, error);
};

const togglePlay = () => {
  if (!sound) return;

  isPlaying.value ? sound.pause() : sound.play();
};

const updateProgress = () => {
  if (sound && isPlaying.value) {
    currentTime.value = sound.seek();
    progress.value = (currentTime.value / duration.value) * 100;
    requestAnimationFrame(updateProgress);
  }
};

const seek = (event) => {
  if (!sound) return;

  const container = event.currentTarget;
  const percent = event.offsetX / container.offsetWidth;
  const seekTime = percent * duration.value;

  sound.seek(seekTime);
  currentTime.value = seekTime;
  progress.value = percent * 100;
};

const updateVolume = () => {
  if (sound) {
    sound.volume(volume.value);
  }
};

const formatTime = (seconds) => {
  if (!seconds || isNaN(seconds)) return '0:00';

  const mins = Math.floor(seconds / 60);
  const secs = Math.floor(seconds % 60);
  return `${mins}:${secs < 10 ? '0' + secs : secs}`;
};

watch(() => props.src, initializeAudio);

onMounted(initializeAudio);

onBeforeUnmount(() => {
  if (sound) {
    sound.unload();
    sound = null;
  }
});
</script>

<style scoped>
.audio-player {
  width: 100%;
  background: linear-gradient(to bottom, #2c2c2c, #1a1a1a);
  border-radius: 8px;
  padding: 15px;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2);
  color: #fff;
}

.controls {
  display: flex;
  align-items: center;
  gap: 15px;
}

.play-button {
  background: transparent;
  border: none;
  cursor: pointer;
  font-size: 24px;
  width: 48px;
  height: 48px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;
  color: #409eff;
  flex-shrink: 0;
}

.play-button:hover {
  color: #66b1ff;
  transform: scale(1.1);
}

.play-button:active {
  transform: scale(0.95);
}

.play-button .el-icon {
  font-size: 28px;
}

.time-progress-container {
  display: flex;
  align-items: center;
  flex-grow: 1;
  gap: 10px;
}

.progress-container {
  flex-grow: 1;
  height: 6px;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 3px;
  cursor: pointer;
  position: relative;
  overflow: visible;
}

.progress-bg {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  border-radius: 3px;
}

.progress-bar {
  height: 100%;
  background: linear-gradient(90deg, #409eff, #66b1ff);
  border-radius: 3px;
  position: relative;
}

.progress-handle {
  position: absolute;
  right: -6px;
  top: 50%;
  transform: translateY(-50%);
  width: 12px;
  height: 12px;
  background: #fff;
  border-radius: 50%;
  box-shadow: 0 0 5px rgba(0, 0, 0, 0.3);
  display: none;
}

.progress-container:hover .progress-handle {
  display: block;
}

.time {
  font-size: 12px;
  color: rgba(255, 255, 255, 0.8);
  font-weight: 500;
  width: 45px;
  text-align: center;
}

.current-time {
  text-align: right;
}

.total-time {
  text-align: left;
}

.volume-container {
  display: flex;
  align-items: center;
  gap: 8px;
  background: rgba(255, 255, 255, 0.1);
  padding: 5px 10px;
  border-radius: 20px;
}

.volume-container .el-icon {
  color: #fff;
  font-size: 16px;
}

.volume-slider-container {
  position: relative;
  width: 60px;
  display: flex;
  align-items: center;
  height: 20px;
}

.volume-slider {
  width: 100%;
  height: 3px;
  border-radius: 1.5px;
  background: rgba(255, 255, 255, 0.2);
  outline: none;
  margin: 0;
  padding: 0;
}

.volume-slider::-webkit-slider-thumb {
  -webkit-appearance: none;
  appearance: none;
  width: 0;
  /* 设置宽度为0，去掉圆圈 */
  height: 0;
  /* 设置高度为0，去掉圆圈 */
  border-radius: 50%;
  background: transparent;
  /* 设置背景为透明 */
  cursor: pointer;
  box-shadow: none;
  /* 去掉阴影 */
}

.volume-slider::-moz-range-thumb {
  width: 0;
  height: 0;
  border-radius: 50%;
  background: transparent;
  box-shadow: none;
}

@media (max-width: 600px) {
  .controls {
    flex-direction: column;
    gap: 10px;
  }

  .time-progress-container {
    width: 100%;
  }

  .volume-container {
    align-self: flex-end;
  }
}
</style>