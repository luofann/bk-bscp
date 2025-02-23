<template>
  <bk-popover
    ref="buttonRef"
    theme="light batch-operation-button-popover"
    placement="bottom-end"
    trigger="click"
    width="108"
    :arrow="false"
    @after-show="isPopoverOpen = true"
    @after-hidden="isPopoverOpen = false">
    <bk-button
      :disabled="props.configs.length === 0 && !props.isAcrossChecked"
      :class="['batch-set-btn', { 'popover-open': isPopoverOpen }]">
      {{ t('批量操作') }}
      <AngleDown class="angle-icon" />
    </bk-button>
    <template #content>
      <div class="operation-item" @click="handleOpenAddToDialog">
        {{ t('批量添加至') }}
      </div>
      <div class="operation-item" @click="handleOpenBantchEditPerm">
        {{ t('批量修改权限') }}
      </div>
      <div v-if="props.pkgType === 'without'" class="operation-item" @click="handleOpenBantchDelet">
        {{ t('批量删除') }}
      </div>
      <div v-if="props.pkgType === 'pkg'" class="operation-item" @click="handleOpenBatchMoveOut">
        {{ t('批量移出套餐') }}
      </div>
    </template>
  </bk-popover>
  <AddToDialog
    v-model:show="isAddToDialogShow"
    :value="props.configs"
    :is-across-checked="props.isAcrossChecked"
    :data-count="props.dataCount"
    @added="emits('refresh')" />
  <EditPermissionDialg
    v-model:show="isEditPermissionShow"
    :loading="editLoading"
    :configs-length="props.isAcrossChecked ? props.dataCount - props.configs.length : props.configs.length"
    :configs="props.configs"
    @confirm="handleConfirmEditPermission($event)" />
  <BatchMoveOutFromPkgDialog
    v-model:show="isBatchMoveDialogShow"
    :current-pkg="props.currentPkg as number"
    :value="props.configs"
    :value-length="props.isAcrossChecked ? props.dataCount - props.configs.length : props.configs.length"
    :is-across-checked="props.isAcrossChecked"
    @moved-out="emits('movedOut')" />
  <MoveOutFromPkgsDialog
    v-model:show="isSingleMoveDialogShow"
    :id="props.configs.length > 0 ? props.configs[0].id : 0"
    :name="props.configs.length > 0 ? props.configs[0].spec.name : ''"
    :current-pkg="props.currentPkg as number"
    @moved-out="emits('movedOut')" />
  <DeleteConfigDialog
    v-model:show="isDeleteConfigDialogShow"
    :is-batch-delete="true"
    :configs="props.configs"
    :is-across-checked="props.isAcrossChecked"
    :data-count="props.dataCount"
    @deleted="emits('deleted')" />
</template>
<script lang="ts" setup>
  import { ref } from 'vue';
  import { AngleDown } from 'bkui-vue/lib/icon';
  import { useI18n } from 'vue-i18n';
  import { useRoute } from 'vue-router';
  import Message from 'bkui-vue/lib/message';
  import { ITemplateConfigItem } from '../../../../../../../../types/template';
  import { batchEditTemplatePermission } from '../../../../../../../api/template';
  import AddToDialog from '../add-to-pkgs/add-to-dialog.vue';
  import EditPermissionDialg from '../edit-permission/edit-permission-dialog.vue';
  import BatchMoveOutFromPkgDialog from '../move-out-from-pkg/batch-move-out-from-pkg-dialog.vue';
  import MoveOutFromPkgsDialog from '../move-out-from-pkg/move-out-from-pkgs-dialog.vue';
  import DeleteConfigDialog from '../delete-configs/delete-config-dialog.vue';

  interface IPermissionType {
    privilege: string;
    user: string;
    user_group: string;
  }

  const { t } = useI18n();
  const route = useRoute();

  const props = defineProps<{
    spaceId: string;
    currentTemplateSpace: number;
    configs: ITemplateConfigItem[];
    pkgType: string;
    currentPkg?: number | string;
    isAcrossChecked: boolean;
    dataCount: number;
  }>();

  const emits = defineEmits(['refresh', 'movedOut', 'deleted']);
  const bkBizId = ref(String(route.params.spaceId));

  const isDeleteConfigDialogShow = ref(false);
  const isEditPermissionShow = ref(false);
  const isPopoverOpen = ref(false);
  const buttonRef = ref();
  const isAddToDialogShow = ref(false);
  const isBatchMoveDialogShow = ref(false);
  const isSingleMoveDialogShow = ref(false);
  const editLoading = ref(false);

  const handleOpenAddToDialog = () => {
    buttonRef.value.hide();
    isAddToDialogShow.value = true;
  };

  const handleOpenBantchEditPerm = () => {
    buttonRef.value.hide();
    isEditPermissionShow.value = true;
  };

  const handleOpenBantchDelet = () => {
    buttonRef.value.hide();
    isDeleteConfigDialogShow.value = true;
  };

  const handleOpenBatchMoveOut = () => {
    buttonRef.value.hide();
    if (props.configs.length === 1) {
      isSingleMoveDialogShow.value = true;
    } else {
      isBatchMoveDialogShow.value = true;
    }
  };

  const handleConfirmEditPermission = async ({
    permission,
    appIds,
  }: {
    permission: IPermissionType;
    appIds: number[];
  }) => {
    try {
      editLoading.value = true;
      const { privilege, user, user_group } = permission;
      const query = {
        privilege,
        user,
        user_group,
        template_ids: props.configs.map((item) => item.id),
        app_ids: appIds,
        template_space_id: props.currentTemplateSpace,
        exclusion_operation: props.isAcrossChecked,
        template_set_id: props.currentPkg !== undefined ? props.currentPkg : 0,
        no_set_specified: props.pkgType === 'without',
      };
      await batchEditTemplatePermission(bkBizId.value, query);
      Message({
        theme: 'success',
        message: t('配置文件权限批量修改成功'),
      });
      isEditPermissionShow.value = false;
      setTimeout(() => {
        emits('refresh');
      }, 300);
    } catch (error) {
      console.error(error);
    } finally {
      editLoading.value = false;
    }
  };
</script>

<style lang="scss" scoped>
  .batch-set-btn {
    min-width: 108px;
    height: 32px;
    margin-left: 8px;
    &.popover-open {
      .angle-icon {
        transform: rotate(-180deg);
      }
    }
    .angle-icon {
      font-size: 20px;
      transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1);
    }
  }
</style>

<style lang="scss">
  .batch-operation-button-popover.bk-popover.bk-pop2-content {
    padding: 4px 0;
    border: 1px solid #dcdee5;
    width: auto !important;
    box-shadow: 0 2px 6px 0 #0000001a;
    .operation-item {
      padding: 0 12px;
      min-width: 58px;
      height: 32px;
      line-height: 32px;
      color: #63656e;
      font-size: 12px;
      cursor: pointer;
      &:hover {
        background: #f5f7fa;
      }
    }
  }
</style>
