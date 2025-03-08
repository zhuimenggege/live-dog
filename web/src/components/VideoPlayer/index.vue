<template>
  <div class="video-player-container" ref="playerContainer">
    <video v-if="props.type === 'flv'" ref="videoElement" controls></video>
    <div v-else ref="mp4Container"></div>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount, watch } from 'vue';
import Player from 'xgplayer';
import Mp4Plugin from 'xgplayer-mp4';
import mpegts from 'mpegts.js';
import 'xgplayer/dist/index.min.css';
import { getToken } from '@/utils/auth';

const props = defineProps({
  src: {
    type: String,
    required: true
  },
  type: {
    type: String,
    default: 'mp4'
  },
  autoplay: {
    type: Boolean,
    default: false
  },
  controls: {
    type: Boolean,
    default: true
  }
});

const playerContainer = ref(null);
const videoElement = ref(null);
const mp4Container = ref(null);
let player = null;
let flvPlayer = null;

const initializePlayer = () => {
  if (player) {
    player.destroy();
    player = null;
  }
  if (flvPlayer) {
    flvPlayer.destroy();
    flvPlayer = null;
  }

  if (props.type === 'flv' && mpegts.isSupported()) {
    const dataSource = {
      type: 'flv',
      url: props.src,
      cors: true,
      hasAudio: true,
      hasVideo: true,
      isLive: false,
    };

    const flvConfig = {
      autoCleanupSourceBuffer: true,
      autoCleanupMaxBackwardDuration: 60,
      autoCleanupMinBackwardDuration: 30,
      enableStashBuffer: true,
      fixAudioTimestampBase: true,
      headers: {
        'Authorization': 'Bearer ' + getToken()
      },
      lazyLoad: true,
      lazyLoadMaxDuration: 60,
      rangeLoadZeroStart: true,
      maxBufferDuration: 30,
      seekType: 'range',
    };

    flvPlayer = mpegts.createPlayer(dataSource, flvConfig);
    flvPlayer.attachMediaElement(videoElement.value);
    flvPlayer.load();

    videoElement.value.addEventListener('seeking', () => {
      if (videoElement.value.currentTime >= videoElement.value.buffered.end(0)) {
        videoElement.value.currentTime = videoElement.value.buffered.end(0) - 1;
      }
    });

    flvPlayer.on(mpegts.Events.ERROR, (errorType, errorDetail) => {
      console.error('MPEGTS Error:', errorType, errorDetail);
    });

    flvPlayer.on(mpegts.Events.BUFFERING, () => {
      console.log('Buffering...');
    });

    flvPlayer.on(mpegts.Events.BUFFER_FULL, () => {
      console.log('Buffer full, ready to play');
    });

    if (props.autoplay) {
      flvPlayer.play().catch(error => {
        console.error('自动播放失败:', error);
      });
    }
  } else {
    player = new Player({
      el: mp4Container.value,
      url: props.src,
      plugins: [Mp4Plugin],
      autoplay: props.autoplay,
      controls: props.controls,
      videoInit: true,
      volume: 1,
      isLive: false,
      preloadTime: 30,
      cors: true,
      responseType: 'blob',
      lang: 'zh-cn',
      mp4plugin: {
        maxBufferLength: 30,
        minBufferLength: 10,
        reqOptions: {
          mode: 'cors',
          headers: {
            'Authorization': 'Bearer ' + getToken(),
          },
        }
      },
    });

    player.on('error', (error) => {
      console.error('播放器错误:', error);
    });

    player.on('ready', () => {
      console.log('播放器已就绪');
    });
  }
};

watch(() => props.src, (newVal) => {
  if ((playerContainer.value || videoElement.value) && newVal) {
    initializePlayer();
  }
});

onMounted(() => {
  if (props.src) {
    initializePlayer();
  }
});

onBeforeUnmount(() => {
  if (flvPlayer) {
    flvPlayer.destroy();
    flvPlayer = null;
  }
  if (player) {
    player.destroy();
    player = null;
  }
});
</script>

<style scoped>
.video-player-container {
  width: 100%;
  margin: 0 auto;
  aspect-ratio: 16 / 9;
  background-color: #000;
}

video {
  width: 100%;
  height: 100%;
}

:deep(.xgplayer) {
  width: 100% !important;
  height: 100% !important;
}

:deep(.xgplayer video) {
  width: 100% !important;
  height: 100% !important;
  object-fit: contain;
}
</style>