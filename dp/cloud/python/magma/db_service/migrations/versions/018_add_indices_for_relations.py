"""empty message

Revision ID: 48e8b58fcc24
Revises: 98f7ccfbd2f8
Create Date: 2022-08-16 13:58:03.352147

"""
import sqlalchemy as sa
from alembic import op

# revision identifiers, used by Alembic.
revision = '48e8b58fcc24'
down_revision = '98f7ccfbd2f8'
branch_labels = None
depends_on = None


def upgrade():
    # ### commands auto generated by Alembic - please adjust! ###
    op.create_index(op.f('ix_channels_cbsd_id'), 'channels', ['cbsd_id'], unique=False)
    op.create_index(op.f('ix_requests_cbsd_id'), 'requests', ['cbsd_id'], unique=False)
    op.create_index(op.f('ix_grants_cbsd_id'), 'grants', ['cbsd_id'], unique=False)
    # ### end Alembic commands ###


def downgrade():
    # ### commands auto generated by Alembic - please adjust! ###
    op.drop_index(op.f('ix_channels_cbsd_id'), table_name='channels')
    op.drop_index(op.f('ix_requests_cbsd_id'), table_name='requests')
    op.drop_index(op.f('ix_grants_cbsd_id'), table_name='grants')
    # ### end Alembic commands ###
