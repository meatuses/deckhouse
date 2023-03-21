import type { Meta, Story } from "@storybook/vue3";

import BaseLayout from "@/components/layout/BaseLayout.vue";
import NodeListPage from "@/pages/NodeListPage.vue";
import { routerDecorator } from "../common";

export default {
  title: "Deckhouse UI/Pages/Node/List",
  component: NodeListPage,
  parameters: { layout: "fullscreen", router: { currentRoute: { name: "NodeListAll" } } },
  decorators: [routerDecorator],
} as Meta;

const Template: Story = (args, { loaded: { releases } }) => ({
  components: { NodeListPage, BaseLayout },

  setup() {
    return { args };
  },

  template: `
    <BaseLayout :compact="false">
      <NodeListPage/>
    </BaseLayout>
  `,
});

export const Default = Template.bind({});
